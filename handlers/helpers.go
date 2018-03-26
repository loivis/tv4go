package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/theckman/go-flock"
)

func ValidateQuery(q url.Values) error {
	required := []string{
		"site",
		"video_id",
		"user_id",
	}
	var missing []string

	for _, v := range required {
		if q.Get(v) == "" {
			missing = append(missing, v)
		}
	}
	log.Printf("missing: %v", missing)
	if len(missing) != 0 {
		return fmt.Errorf("missing required query parameters: %v", missing)
	}

	return nil
}

func LockFile(userID string) (*flock.Flock, error) {
	filePath := path.Join(os.TempDir(), userID+".info")
	fileLock := flock.NewFlock(filePath)

	lockCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, 500*time.Millisecond)

	if err != nil {
		return nil, err
	}

	if !locked {
		return nil, fmt.Errorf("failed to get file lock on user %s", userID)
	}

	fmt.Printf("path: %s; locked: %v\n", fileLock.Path(), fileLock.Locked())

	return fileLock, nil
}

func UpdateUserInfo(path, action, videoID string) error {
	info := Info{Active: make(map[string]bool)}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err)
	}
	json.Unmarshal(content, &info)

	if action == "start" {
		if len(info.Active) >= 2 {
			return fmt.Errorf("users can only watch 2 video streams at the same time: %v", info.Active)
		}
		info.Active[videoID] = true
	}
	if action == "stop" {
		if !info.Active[videoID] {
			return fmt.Errorf("you're not watching video %s", videoID)
		}
		delete(info.Active, videoID)
	}

	info.Events = append(info.Events, &Event{
		Time:    time.Now(),
		VideoID: videoID,
		Action:  action,
	})
	log.Printf("info: %v", info)
	content, err = json.Marshal(info)
	if err != nil {
		return fmt.Errorf("error updating user info: %s", err)
	}
	err = ioutil.WriteFile(path, content, 644)
	if err != nil {
		return fmt.Errorf("error updating user info: %s", err)
	}

	return nil
}

func formatResponse(response *response, status, msg string) []byte {
	log.Println(msg)
	response.Status = status
	response.Message = msg
	resp, _ := json.Marshal(response)
	return resp
}
