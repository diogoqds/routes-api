package config

import (
	"encoding/json"
	"github.com/diogoqds/routes-challenge-api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ConfigHttpRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(map[string]interface{}{"hello": "world"}); err != nil {
			log.Println("[Response Body Error]", err)
			return
		}

	}).Methods(http.MethodGet)

	router.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
}
