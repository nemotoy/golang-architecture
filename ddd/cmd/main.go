package main

import (
	"log"
	"net/http"
)

func main() {
	runServer()
}

func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Printf("%v", err)
	}
}

func helloHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}
}
