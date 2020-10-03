package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var event1 = event{
	ID:          "1",
	Title:       "first event",
	Description: "Very first event",
}

var events = allEvents{event1}

func getEvents(w http.ResponseWriter, r *http.Request) {
	eventsByte, err := json.Marshal(events)
	if err != nil {

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	http.ListenAndServe(":8181", router)
}
