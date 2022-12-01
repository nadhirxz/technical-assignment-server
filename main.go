package main

import (
	"log"

	"github.com/joho/godotenv"
	"gitlab.com/nadhirxz/technical-assignment-server/server"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Errr loading .env file")
	}

	server.Init()
}
