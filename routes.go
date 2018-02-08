package main

import (
	"net/http"

    "github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Path string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)		

		router.Methods(route.Method).Path(route.Path).Name(route.Name).Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"IndexHandler",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"TodoIndexHandler",
		"GET",
		"/todos",
		TodoIndexHandler,
	},
	Route{
		"TodoShowHandler",
		"GET",
		"/todos/{todoId}",
		TodoShowHandler,
	},
}
