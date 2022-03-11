package main

import (
	"github.com/mserebryaakov/books-rest-api/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()

	s := apiserver.New(config)
	err := s.Start()
	if err != nil {
		config.Log.Fatal(err)
	}
}
