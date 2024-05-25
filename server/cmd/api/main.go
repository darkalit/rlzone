package main

import (
	"log"

	"github.com/darkalit/rlzone/server/http"
)

func main() {
	s, err := http.NewServer()
	if err != nil {
		log.Fatal("Failed to create server")
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
