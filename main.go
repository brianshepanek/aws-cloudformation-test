package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/fizz", fizzHandler)
	fmt.Println("Server is listening on port 8080...")
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
