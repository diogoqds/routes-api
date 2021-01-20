package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/diogoqds/routes-challenge-api/usecases"
)

type requestParams struct {
	Email string `json:"email"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var bodyParams requestParams
	if err = json.Unmarshal(body, &bodyParams); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := usecases.Authenticate(bodyParams.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		log.Println("[JWT Error]", err)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"token": token}); err != nil {
		log.Println("[Response Body Error]", err)
		return
	}
}
