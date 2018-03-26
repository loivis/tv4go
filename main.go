package main

import (
	"net/http"

	"github.com/TV4/graceful"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router(),
	}
	graceful.LogListenAndServe(srv)
}
