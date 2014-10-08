package main

import (
	"log"
	"net/http"
)

func fooServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello\n"))
}

func main() {
	http.HandleFunc("/foo", fooServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
