package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

// var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	/*
		The list of birds is now taken from the store instead of the package level  `birds` variable we had earlier
		The `store` variable is the package level variable that we defined in
		`store.go`, and is initialized during the initialization phase of the
		application
	*/
	birds, err := store.GetBirds()

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

	// Add the new bird details in to store
	err = store.CreateBird(&bird)
	if err != nil {
		fmt.Println(err)
	}

	// Finally redirect the user to original HTML page using the httm library 'Redirect' method
	http.Redirect(w, r, "http://localhost:8080/assets/web_application/assets/", http.StatusFound)
}
