package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Bird struct {
	Species string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the bird variable to json
	birdListBytes, err := json.Marshal(birds)
	// If there is an error, print it to console, and return a server error response to user
	if err != nil {
		fmt.Println(fmt.Errorf("Error:%v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well write the JSON list of birds to response
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of bird
	bird := Bird{}

	// We send all our data in html form data, the 'ParseForm' method of the request parses the form values
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from form info
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// Append the existing list of birds with new entry
	birds = append(birds, bird)

	// Finally redirect the user to original HTML page using the httm library 'Redirect' method
	http.Redirect(w, r, "http://localhost:8080/assets/web_application/assets/", http.StatusFound)
}