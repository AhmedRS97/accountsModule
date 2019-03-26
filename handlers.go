package main

import (
	"encoding/json"
	//"fmt"
	"net/http"
        //"github.com/jinzhu/gorm"
	//"github.com/gorilla/mux"
)

func AccountSignup(w http.ResponseWriter, r *http.Request) {
	// the controller for the account sign up route

	// this will make a Account model object to store the json data in it 
	account := &Account{}

	//decode the request body into struct and failed if any error occur
	err := json.NewDecoder(r.Body).Decode(account) 
	if err != nil {
		Respond(w, Message(false, "Invalid request"))
		return
	}

	// return the response for the client
        Respond(w, account.Create())
}

/*func EmailConfirmation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode('"this endpoint will be implemented later"'); err != nil {
		panic(err)
	}
}*/
