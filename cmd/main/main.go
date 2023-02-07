package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	storyutils "github.com/humberto1212/gopher-story/pkg/storyUtils"
	"github.com/humberto1212/gopher-story/pkg/utils"
)

func main() {

	history := utils.ParseJSON()

	fmt.Println(history)

	r := mux.NewRouter()
	r.HandleFunc("/intro", storyutils.IntroHandler).Methods("GET")
	r.HandleFunc("/new-york", storyutils.NewYorkHandler).Methods("GET")
	r.HandleFunc("/sean-kelly", storyutils.SeanKellyHandler).Methods("GET")
	r.HandleFunc("/mark-bates", storyutils.MarkBatesHandler).Methods("GET")
	r.HandleFunc("/denver", storyutils.DenverHandler).Methods("GET")
	r.HandleFunc("/home", storyutils.HomeHandler).Methods("GET")

	// Start the HTTP server
	http.ListenAndServe(":8000", r)
}
