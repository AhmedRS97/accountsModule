package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	// this function will make a router for each handler
	// it's using the mux lib for the implementation of the router

	router := mux.NewRouter().StrictSlash(true)
	// here it will loop over the routes from the routes.go file
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		// here will make a logger for the handler
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
