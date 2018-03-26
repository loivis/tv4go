package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/loivis/tv4go/handlers"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}
	rootDir := path.Dir(filename)
	fmt.Println(rootDir)
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.Health).Methods("GET")
	r.HandleFunc("/startstream", handlers.Startstream).Methods("GET")
	r.HandleFunc("/stopstream", handlers.Stopstream).Methods("GET")
	return r
}
