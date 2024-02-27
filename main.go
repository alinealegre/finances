package main

import (
	"finances/database"
	"finances/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.StartDB()
	s := server.NewServer()

	s.Run()
}
