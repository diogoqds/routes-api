package config

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/diogoqds/routes-challenge-api/controllers"
	"github.com/diogoqds/routes-challenge-api/middlewares"
	"github.com/gorilla/mux"
)

func ConfigHttpRoutes(router *mux.Router) {
	router.Handle("/", middlewares.AuthMiddleware(Hello))
	router.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	router.HandleFunc("/sellers", middlewares.AuthMiddleware(controllers.Seller.Create)).Methods(http.MethodPost)
	router.HandleFunc("/sellers", middlewares.AuthMiddleware(controllers.Seller.Index)).Methods(http.MethodGet)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"hello": "world"}); err != nil {
		log.Println("[Response Body Error]", err)
		return
	}
}
