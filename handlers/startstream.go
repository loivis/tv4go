package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	search "github.com/TV4/search-go"
)

func Startstream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response response

	log.Println(r.RequestURI, r.Method, r.Header["User-Agent"])
	queryIn := r.URL.Query()
	err := ValidateQuery(queryIn)
	if err != nil {
		w.Write(formatResponse(&response, "error", err.Error()))
		return
	}
	site := queryIn.Get("site")
	videoID := queryIn.Get("video_id")
	userID := queryIn.Get("user_id")
	response.VideoID = videoID

	file, err := LockFile(userID)
	if err != nil {
		w.Write(formatResponse(&response, "error", err.Error()))
		return
	}
	defer file.Unlock()

	client := search.NewClient(
		search.SetDebugLogf(log.New(os.Stderr, "", 0).Printf),
	)
	queryOut := url.Values{
		"site":      {site},
		"video_ids": {videoID},
	}
	log.Printf("perform search: %v", queryOut)
	hits, err := client.Search(context.Background(), queryOut)
	if err != nil {
		w.Write(formatResponse(&response, "error", err.Error()))
		return
	}

	if hits.TotalHits != 1 {
		msg := fmt.Sprintf("video with id %s not found", videoID)
		w.Write(formatResponse(&response, "error", msg))
		return
	}

	err = UpdateUserInfo(file.Path(), "start", videoID)
	if err != nil {
		w.Write(formatResponse(&response, "error", err.Error()))
		return
	}

	msg := fmt.Sprintf("video with id %s found: %s", videoID, hits.Hits[0].Subset().TitleSv)
	w.Write(formatResponse(&response, "started", msg))
}
