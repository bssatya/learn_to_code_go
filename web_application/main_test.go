// This is the name of our package

// Everything with this pacakge name can see everything

// else inside the same package, regardless of the file they are in

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handler(t *testing.T) {
	// Here we form a new HTTP request. This is the request that's going to be passed to our handler.
	// The first argument is method, the second argument is the route (which we have blank for now, and will get back to soon),
	// and the third is the request body, which we don't have in this case.
	req, err := http.NewRequest("GET", "", nil)
	// In case there is a error forming the request, we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	// We use go's httptest library to create an http recorder. This recorder will act as the target of our http request
	// (think of it as a mini brower which accept the result of the http request we make)
	recorder := httptest.NewRecorder()

	// Create an HTTP handler from our handler function. "handler" is the handler function defined in main.go file that
	// we want to test
	hf := http.HandlerFunc(handler)

	// Serve the HTTP request to our recorder. This is the line that actually executes handler that we want to test
	hf.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := "Hello World\n"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRouteForExistantRoute_newRouter(t *testing.T) {
	// Instantiate the router using the constructor function that we defined previously
	r := newRouter()

	// Create a new server using the "httptest" libraries 'NewServer' method
	mockServer := httptest.NewServer(r)

	// The mock server created runs the server and exposes its location in the URL attribute
	// We make get request to the 'hello' route that we defined in the router
	resp, err := http.Get(mockServer.URL + "/hello")

	// handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string
	defer resp.Body.Close()
	// Read the body in to a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Converts the bytes to string
	respString := string(b)
	expected := "Hello World\n"

	// The response should match with the one defined in our handler.
	// If it does happen to be "Hello World\n", then it confirms, that the route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouteForNonExistantRoute_newRouter(t *testing.T) {
	// Instantiate the router using the constructor function that we defined previously
	r := newRouter()

	// Create a new server using the "httptest" libraries 'NewServer' method
	mockServer := httptest.NewServer(r)

	// The mock server created runs the server and exposes its location in the URL attribute
	// We make get request to the 'hello' route that we defined in the router
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	// handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string
	defer resp.Body.Close()
	// Read the body in to a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Converts the bytes to string
	respString := string(b)
	expected := ""

	// The response should match with the one defined in our handler.
	// If it does happen to be "Hello World\n", then it confirms, that the route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// Hit the 'GET /assets/' route to get the index.html file response
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	// The status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// It isn't wise to test the entire content of the html file.
	// Instead, we test the content-type header is "text/html"; charset=utf-8"
	// So that we know that the html file has been served
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}
