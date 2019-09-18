package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nemotoy/golang-architecture/ddd/config"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	runServer(conf)
}

func runServer(conf *config.Config) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler())

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), mux); err != nil {
		log.Printf("%v", err)
	}
}

func helloHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}
}
