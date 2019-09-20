package handler

import "net/http"

type GreetHandler interface {
	Greet(w http.ResponseWriter, r *http.Request)
}

type greetHandler struct {
}

func NewGreetHandler() GreetHandler {
	return &greetHandler{}
}

func (g *greetHandler) Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
