package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetBirdHandler(t *testing.T) {
	// birds := make([]Bird, 3)
	// birds[0] = Bird{
	// 	Species:     "sparow",
	// 	Description: "A small harmless bird",
	// }
	// Initialize the mock store
	mockStore := InitMockStore()

	/*
		Define the data that we want to return when the mocks 'Get Birds' method is called. Also, we expect
		it to be called only once.
	*/
	mockStore.On("GetBirds").Return([]*Bird{{"sparow", "A small harmless bird"}}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getBirdHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Bird{"sparow", "A small harmless bird"}
	b := []Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body, got %v want %v", actual, expected)
	}

	// the expectations that we defined in the `on` method are asserted here
	mockStore.AssertExpectations(t)
}

func TestCreateBirdsHandler(t *testing.T) {
	// birds := make([]Bird, 3)
	// birds[0] = Bird{
	// 	Species:     "sparow",
	// 	Description: "A small harmless bird",
	// }
	mockStore := InitMockStore()
	/*
	 Similarly, we define our expectations for th `CreateBird` method.
	 We expect the first argument to the method to be the bird struct
	 defined below, and tell the mock to return a `nil` error
	*/
	mockStore.On("CreateBird", &Bird{"eagle", "A bird of prey"}).Return(nil)

	form := newCreateBirdForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createBirdHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// expected := Bird{"eagle", "A bird of Prey"}

	// actual := birds[1]

	// if actual != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	// }
	mockStore.AssertExpectations(t)
}

func newCreateBirdForm() *url.Values {
	form := url.Values{}
	form.Set("species", "eagle")
	form.Set("description", "A bird of Prey")
	return &form
}
