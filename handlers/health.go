package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI, r.Method, r.Header["User-Agent"])
	// logic to relect real health status
	response, _ := json.Marshal(map[string]string{"status": "ok"})
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
