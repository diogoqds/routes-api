package main

import (
	"github.com/diogoqds/routes-challenge-api/config"
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	log.Println("Starting the application")
	infra.SetupDB()
	router := infra.CreateRouter()
	config.ConfigHttpRoutes(router)
	log.Println("Server running on ", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}
