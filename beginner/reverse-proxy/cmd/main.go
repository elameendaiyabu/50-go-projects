package main

import (
	"go-projects/beginner/reverse-proxy/internal/server"
	"log"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
