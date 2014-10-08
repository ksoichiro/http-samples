package main

import (
	"log"
	"net/http"
)

// fooServer func has http.Handler interface
func fooServer(w http.ResponseWriter, r *http.Request) {
	// Set HTTP status code by myself(optional)
	w.WriteHeader(http.StatusOK)
	// Write response body
	w.Write([]byte("Hello\n"))
}

func main() {
	// Handle path '/foo' by fooServer func
	http.HandleFunc("/foo", fooServer)
	// Serve with port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
