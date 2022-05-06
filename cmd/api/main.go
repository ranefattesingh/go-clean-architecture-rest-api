package main

import (
	"log"

	"github.com/ranefattesingh/go-clean-architecture-rest-api/internal/server"
)

func main() {

	s := server.NewServer()
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
