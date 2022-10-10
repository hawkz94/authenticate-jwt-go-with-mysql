package main

import (
	"log"
	"tecmentor-api/database"
	"tecmentor-api/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
