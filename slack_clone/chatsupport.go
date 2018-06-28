package main // creates an executable

import (
	"fmt"
	"net/http"
)

func main() {
	// pass in URL pattern and handler fcn to execute
	http.HandleFunc("/", handler)
	// waits until port is connected to or app is blocked
	http.ListenAndServe(":4000", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	// pass in a writer and string we'd like written as a response
	fmt.Fprintf(w, "this is how we create a web server in Go")
}
