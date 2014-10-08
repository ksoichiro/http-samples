package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// fooServer func has http.Handler interface
func fooServer(w http.ResponseWriter, r *http.Request) {
	// Set HTTP status code by myself(optional)
	w.WriteHeader(http.StatusOK)
	// Write response body
	w.Write([]byte("Hello, world!\n"))
}

func main() {
	portNum := flag.Int("port", 8080, "HTTP server port")
	flag.Parse()

	port := fmt.Sprintf("%d", *portNum)

	// Handle path '/foo' by fooServer func
	http.HandleFunc("/hello", fooServer)

	// Serve with port var port
	fmt.Printf("Serves HTTP in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
