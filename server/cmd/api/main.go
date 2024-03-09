package main

import (
	"log"

	"github.com/darkalit/rlzone/server/http/rest"
)

func main() {
	s, err := rest.NewServer()
	if err != nil {
		log.Fatal("Failed to create server")
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
