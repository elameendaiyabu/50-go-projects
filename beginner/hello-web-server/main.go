package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
