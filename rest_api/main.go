package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	err := json.NewEncoder(w).Encode(events)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}

func errWriter(w http.ResponseWriter, statusCode int, errMsg string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(errMsg))
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errWriter(w, http.StatusInternalServerError, err.Error())
	}
	err = json.Unmarshal(reqBody, &newEvent)
	if err != nil {
		errWriter(w, http.StatusInternalServerError, err.Error())
	}
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	var updatedEvent event
	id := mux.Vars(r)["id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errWriter(w, http.StatusBadRequest, err.Error())
	}
	err = json.Unmarshal(reqBody, &updatedEvent)
	if err != nil {
		errWriter(w, http.StatusBadRequest, err.Error())
	}
	eventFound := false
	for _, e := range events {
		if e.ID == id {
			e.Description = updatedEvent.Description
			e.Title = updatedEvent.Title
			eventFound = true
		}
	}
	if !eventFound {
		errWriter(w, http.StatusBadRequest, "Event ID not found")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedEvent)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/getevents", getEvents)
	router.HandleFunc("/createvent", createEvent)
	router.HandleFunc("/updateEvent", updateEvent)
	http.ListenAndServe(":8181", router)
}
