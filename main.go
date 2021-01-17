package main

import (
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Println("Starting the application")
	infra.SetupDB()
}
