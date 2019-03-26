package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// the routes for the router to handle
	Route{
		"AccountSignup",
		"PUT",
		"/api/v1/auth/signup",
		AccountSignup,
	},
	/*Route{
		"EmailConfirmation",
		"PUT",
		"/api/v1/auth/confirm-email/{uuid}",
		EmailConfirmation,
	},*/
}
