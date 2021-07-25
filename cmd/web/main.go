package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	mux := routes()

	log.Println("Starting web server on " + portNumber)
	_ = http.ListenAndServe(portNumber, mux)
}
