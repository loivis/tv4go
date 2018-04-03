package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Stopstream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response response

	log.Println(r.RequestURI, r.Method, r.Header["User-Agent"])
	queryIn := r.URL.Query()
	err := ValidateQuery(queryIn)
	if err != nil {
		return
	}
	videoID := queryIn.Get("video_id")
	userID := queryIn.Get("user_id")
	response.VideoID = videoID

	file, err := LockFile(userID)
	if err != nil {
		w.Write(formatResponse(&response, "error", err.Error()))
		return
	}
	defer file.Unlock()

	err = UpdateUserInfo(file.Path(), "stop", videoID)
	if err != nil {
		w.Write(formatResponse(&response, "error", err.Error()))
		return
	}

	msg := fmt.Sprintf("video with id %s found: ", videoID)
	w.Write(formatResponse(&response, "stopped", msg))
}
