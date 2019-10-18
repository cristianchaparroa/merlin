package main

import (
	"github.com/cristianchaparroa/merlin/bycles/api"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	s := api.NewServer()
	s.Run()
}
