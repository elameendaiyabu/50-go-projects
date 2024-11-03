package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	HelloWorld(res, req)

	got := res.Body.String()
	want := "Hello world"

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
