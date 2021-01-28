package main

import (
	"log"
	"net/http"

	"github.com/diogoqds/routes-challenge-api/config"
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/joho/godotenv"
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
	router := infra.CreateRouter()
	config.ConfigHttpRoutes(router)
	log.Println("Server running on port", 3000)
	log.Fatal(http.ListenAndServe(":3000", router))
}
