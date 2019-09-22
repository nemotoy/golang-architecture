package handler

import "net/http"

type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (u *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
