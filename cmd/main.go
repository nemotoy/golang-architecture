package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nemotoy/golang-architecture/config"
	"github.com/nemotoy/golang-architecture/interfaces/handler"
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
	uh := handler.NewUserHandler()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) { uh.Get(w, r) })

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), mux); err != nil {
		log.Printf("%v", err)
	}
}
