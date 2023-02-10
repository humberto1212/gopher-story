package main

import (
	"net/http"

	"github.com/gorilla/mux"
	storyutils "github.com/humberto1212/gopher-story/pkg/storyUtils"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", storyutils.IntroHandler).Methods("GET")
	r.HandleFunc("/new-york", storyutils.NewYorkHandler).Methods("GET")
	r.HandleFunc("/debate", storyutils.DebateHandler).Methods("GET")
	r.HandleFunc("/sean-kelly", storyutils.SeanKellyHandler).Methods("GET")
	r.HandleFunc("/mark-bates", storyutils.MarkBatesHandler).Methods("GET")
	r.HandleFunc("/denver", storyutils.DenverHandler).Methods("GET")
	r.HandleFunc("/home", storyutils.HomeHandler).Methods("GET")

	// Start the HTTP server
	http.ListenAndServe(":8000", r)
}
