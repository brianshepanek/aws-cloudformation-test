package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/fizz", fizzHandler)
	http.HandleFunc("/fuzz", fuzzHandler)
	http.HandleFunc("/rand", randHandler)
	fmt.Println("Server is listening on port 80...")
	http.ListenAndServe(":80", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World @ "+time.Now().String())
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Foo @ "+time.Now().String())
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Bar @ "+time.Now().String())
}

func fizzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Fizz @ "+time.Now().String())
}

func fuzzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Fuzz @ "+time.Now().String())
}

func randHandler(w http.ResponseWriter, r *http.Request) {
	// Predefined sentences
	sentences := []string{
		"Hello, world!",
		"The quick brown fox jumps over the lazy dog.",
		"Programming in Go is fun.",
		"AI is the future.",
		"OpenAI has created the GPT model.",
		"This is a sample sentence.",
		"The sun rises in the east.",
		"Be kind to others.",
		"Hard work always pays off.",
		"Success is the result of perseverance.",
		"Honesty is the best policy.",
		"Health is wealth.",
		"The early bird catches the worm.",
		"Practice makes perfect.",
		"Rome was not built in a day.",
		"Unity is strength.",
		"Education is the key to success.",
		"Time and tide wait for no man.",
		"Action speaks louder than words.",
		"All that glitters is not gold.",
		"Where there's smoke, there's fire.",
		"You can't judge a book by its cover.",
		"When in Rome, do as the Romans do.",
		"Every cloud has a silver lining.",
		"Don't count your chickens before they're hatched.",
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Get a random index
	index := rand.Intn(len(sentences))

	// Print the random sentence
	fmt.Fprintf(w, sentences[index]+" @ "+time.Now().String())
}
