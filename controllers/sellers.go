package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/diogoqds/routes-challenge-api/usecases"
	"io/ioutil"
	"log"
	"net/http"
)

type SellerController struct {
}

func (s SellerController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var bodyParams map[string]interface{}

	if err = json.Unmarshal(body, &bodyParams); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := fmt.Sprintf("%s", bodyParams["name"])
	email := fmt.Sprintf("%s", bodyParams["email"])

	seller, err := usecases.CreateSellerService.Create(name, email)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(seller); err != nil {
		log.Println("[Response Body Error]", err)
		return
	}
}

var (
	Seller = SellerController{}
)
