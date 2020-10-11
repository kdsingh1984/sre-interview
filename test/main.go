package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Table struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

func numTable(num int) []string {
	var returnList []string
	for i := 1; i <= 10; i++ {
		r := i * num
		returnList = append(returnList, fmt.Sprintf("%v x %v = %v", num, i, r))
	}
	return returnList
}

func welcome(w http.ResponseWriter, r *http.Request) {
	var t Table
	t.ID = 1
	t.Name = "Kamla"
	// number, _ := strconv.Atoi(mux.Vars(r)["id"])
	// numResultList := strings.Join(numTable(number), "\n")
	json.NewEncoder(w).Encode(t)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{id}", welcome)
	http.ListenAndServe("localhost:8888", http.Handler(router))
}
