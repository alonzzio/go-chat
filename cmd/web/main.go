package main

import (
	"github.com/alonzzio/go-chat/internal/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on " + portNumber)
	_ = http.ListenAndServe(portNumber, mux)
}
