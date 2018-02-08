package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name string
	Completed bool
	Due time.Time
}

type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/todos", TodoIndexHandler)
	router.HandleFunc("/todos/{todoId}", TodoShowHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo index!")	
}

func TodoShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	fmt.Fprintln(w, "Todo show: ", todoId)	
}