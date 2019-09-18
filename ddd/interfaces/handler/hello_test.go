package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(HelloHandler()))
	defer s.Close()

	res, err := http.Get(s.URL)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	want := "Hello, world!"
	if string(body) != want {
		t.Error(err)
	}
}
