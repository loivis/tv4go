package main

import (
	"github.com/loivis/tv4go/handlers"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	srv := handlers.Server{
		Store: &handlers.FileUserInfoStore{},
	}

	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.Health).Methods("GET")
	r.HandleFunc("/startstream", handlers.Startstream).Methods("GET")
	r.HandleFunc("/stopstream", srv.Stopstream).Methods("GET")
	return r
}
