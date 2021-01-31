package controllers

import (
	"github.com/diogoqds/routes-challenge-api/usecases"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type SellerClientController struct {
}

func (s SellerClientController) Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sellerId, err := strconv.Atoi(vars["seller_id"])

	clients, err := usecases.ListSellerClientsService.FindAll(sellerId)

	if err != nil {
		WriteResponse(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	WriteResponse(
		w,
		http.StatusOK,
		map[string]interface{}{"clients": clients},
	)
}

var (
	SellerClient = SellerClientController{}
)
