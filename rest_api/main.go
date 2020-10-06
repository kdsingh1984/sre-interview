package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:MyNewPass@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
	return db
}

var DB *sql.DB = dbConnect()

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func readTags(w http.ResponseWriter, r *http.Request) {
	var tag Tag
	var tags []Tag
	rows, err := DB.Query("select  * from tags")
	if err != nil {
		errWriter(w, http.StatusBadRequest, err.Error())
		return
	}
	for rows.Next() {
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			errWriter(w, 500, err.Error())
			return
		}
		tags = append(tags, tag)
	}
	json.NewEncoder(w).Encode(tags)
}

// func postCall(w http.ResponseWriter, r *http.Request, reqStruct interface{}) interface{} {
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		errWriter(w, http.StatusBadRequest, err.Error())
// 		return nil
// 	}
// 	err = json.Unmarshal(reqBody, &reqStruct)
// 	if err != nil {
// 		errWriter(w, http.StatusInternalServerError, err.Error())
// 		return nil
// 	}
// 	return reqStruct
// }

func addTag(w http.ResponseWriter, r *http.Request) {
	var tag Tag
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errWriter(w, http.StatusBadRequest, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &tag)
	if err != nil {
		errWriter(w, http.StatusInternalServerError, err.Error())
		return
	}
	q := fmt.Sprintf("insert into %s (id,name) values (%d,'%s');", "tags", tag.ID, tag.Name)
	_, err = DB.Exec(q)
	if err != nil {
		errWriter(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	defer DB.Close()
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/getevents", getEvents)
	router.HandleFunc("/createvent", createEvent)
	router.HandleFunc("/updateEvent", updateEvent)
	router.HandleFunc("/getallTags", readTags)
	router.HandleFunc("/addTag", addTag)
	http.ListenAndServe(":8181", router)
}
