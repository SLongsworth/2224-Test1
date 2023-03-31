package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//create a map(multiplexer)
	mux := http.NewServeMux()

	//create handlers for the routers
	mux.HandleFunc("/", myBioHandler)
	//mux.HandleFunc("/ran", randHandler)
	//mux.HandleFunc("/greet", greetingHandler)

	//create a server/port
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}

func myBioHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
		<html>
			<head>
				<title>myBio</title>
				<style>
					body {
						font-size: 30px;
						text-align: center;
				</style>
			</head>
			<body>
			Good day! My name is Shia Teyan Longsworth, I am 19 years old. My birthday is on the
			28th of May.I love to cook meals especially pasta. I enjoy hanging out with friends and spending
			time with my family. I enjoy watching series and movies and rewatching them. I dislike sports,
			waking up early in the mornings, and white condiments.		
			</body>
		</html>
	`
	fmt.Fprint(w, html)
}
