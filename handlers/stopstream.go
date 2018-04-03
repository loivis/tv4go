package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Store UserInfoStore
}

func (srv *Server) Stopstream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response response

	log.Println(r.RequestURI, r.Method, r.Header["User-Agent"])
	queryIn := r.URL.Query()
	err := ValidateQuery(queryIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatResponse(&response, "error", err.Error()))
		return
	}
	videoID := queryIn.Get("video_id")
	userID := queryIn.Get("user_id")
	response.VideoID = videoID

	if err := srv.Store.Stop(userID, videoID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatResponse(&response, "error", err.Error()))
		return
	}

	msg := fmt.Sprintf("video with id %s found: ", videoID)
	w.Write(formatResponse(&response, "stopped", msg))
}
