package main

import (
	"log"
	"net/http"
)

func main() {
	//create a map(multiplexer)
	mux := http.NewServeMux()

	//create handlers for the routers
	mux.HandleFunc("/", myBioHandler)
	mux.HandleFunc("/ran", randHandler)
	mux.HandleFunc("/greet", greetingHandler)

	//create a server/port
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
