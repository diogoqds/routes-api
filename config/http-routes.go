package config

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ConfigHttpRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err:= json.NewEncoder(w).Encode(map[string]interface{}{ "hello": "world" }); err != nil {
			log.Println("[Response Body Error]", err)
			return
		}

	}).Methods(http.MethodGet)
}