package main

import (
	"encoding/json"
	"net/http"
)


func Message(status bool, message string) (map[string]interface{}) {
        // this function makes a message object and return it to the caller
        return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
        // this function makes
        w.Header().Add("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
}

