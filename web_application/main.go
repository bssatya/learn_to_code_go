// This is the name of our package
// Everything with this pacakge name can see everything
// else inside the same package, regardless of the file they are in

package main

// These are the libraies we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	"fmt" // The "net/http" library has methods to implement HTTP clients and servers
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and returns it to us. We can now use this function
// to instantiate and test the router outside of the main function.
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directy and point it to the directry 'assets' we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename. The fileserver is wrapped in the `stripPrefix`
	// method, because we want to remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server will look for only "index.html"
	//  inside the directory declared above. If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting with "/assets/", instead of the
	// absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {
	// Declare a new router
	r := newRouter()

	// We can pass our router (after declaring all the routes) to this method
	// (where previously, we were leaving the second argument as nil)
	http.ListenAndServe(":8080", r)
}

// handler is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World\n")

}
