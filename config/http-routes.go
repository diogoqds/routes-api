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
	router.HandleFunc("/sellers/{id}", middlewares.AuthMiddleware(controllers.Seller.Delete)).Methods(http.MethodDelete)
	router.HandleFunc("/routes", middlewares.AuthMiddleware(controllers.Routes.Create)).Methods(http.MethodPost)
	router.HandleFunc("/routes/{id}", middlewares.AuthMiddleware(controllers.Routes.Update)).Methods(http.MethodPut)
	router.HandleFunc("/routes/{id}", middlewares.AuthMiddleware(controllers.Routes.Delete)).Methods(http.MethodDelete)
	router.HandleFunc("/routes/{id}/seller", middlewares.AuthMiddleware(controllers.Routes.AssociateSeller)).Methods(http.MethodPatch)
	router.HandleFunc("/routes/{id}/seller", middlewares.AuthMiddleware(controllers.Routes.DisassociateSeller)).Methods(http.MethodDelete)
	router.HandleFunc("/clients", controllers.Client.Create).Methods(http.MethodPost)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"hello": "world"}); err != nil {
		log.Println("[Response Body Error]", err)
		return
	}
}
