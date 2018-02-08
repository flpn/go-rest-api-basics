package main

import "net/http"

type Route struct {
	Name string
	Method string
	Path string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
