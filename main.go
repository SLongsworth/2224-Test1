package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	//create a map(multiplexer)
	mux := http.NewServeMux()

	//create handlers for the routers
	mux.HandleFunc("/", myBioHandler)
	mux.HandleFunc("/ran", randoHandler)
	mux.HandleFunc("/greet", greetingHandler)

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

func randoHandler(w http.ResponseWriter, r *http.Request) {
	randQ := make([]string, 0)
	randQ = append(randQ,
		"Everything has beauty but not everyone sees it",
		"You cant cross the sea merely by standing and staring at the water",
		"Im about as intimidating as a butterfly",
		"Whatever you are, be a good one",
		"Be yourself everyone else is already taken",
		"Life shrinks or expands in proportion to ones courage",
		"Nothing is impossible The word itself says Im possible",
		"If you can dream it you can do it")

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(randQ)) + 1

	html := `
	<!DOCTYPE html>
		<html>
			<head>
				<title>Random Quotes</title>
				<style>
					body {
						font-size: 30px;
						text-align: center;
				</style>
			</head>
			<body>
			` + randQ[random] + `	
			</body>
		</html>
	`
	fmt.Fprintf(w, html)
}

// greetingHandler handles the /greeting route
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	nowTime := time.Now()
	greetDay := time.Now().Weekday().String()
	html := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Greeting</title>
				<style>
					body {
						font-size: 30px;
						text-align: center;
				</style>
			</head>
			<body>
				<p1>Right now is {{time}}</p1>
				<p2>Enjoy the rest of your {{greetDay}}.</p2>
			</body>
		</html>
	`
	html = strings.ReplaceAll(html, "{{time}}", nowTime.Format("3:04PM"))
	html = strings.ReplaceAll(html, "{{greetDay}}", greetDay)
	fmt.Fprint(w, html)
}
