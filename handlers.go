package main

import (
    "encoding/json"
    "fmt"
	"net/http"
	"io/ioutil"
	"io"

    "github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!!!!!")	
}

func TodoIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}	
}

func TodoShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	fmt.Fprintln(w, "Todo show: ", todoId)	
}

func TodoCreateHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	newTodo := RepoCreateTodo(todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	
    if err := json.NewEncoder(w).Encode(newTodo); err != nil {
        panic(err)
    }
}
