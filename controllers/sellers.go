package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/diogoqds/routes-challenge-api/usecases"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SellerController struct {
}

func (s SellerController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(
			w,
			http.StatusNotFound,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	var bodyParams map[string]interface{}

	if err = json.Unmarshal(body, &bodyParams); err != nil {
		WriteResponse(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	name := fmt.Sprintf("%s", bodyParams["name"])
	email := fmt.Sprintf("%s", bodyParams["email"])

	seller, err := usecases.CreateSellerService.Create(name, email)

	if err != nil {
		WriteResponse(
			w,
			http.StatusUnprocessableEntity,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	WriteResponse(
		w,
		http.StatusCreated,
		map[string]interface{}{"seller": seller},
	)

}

func (s SellerController) Index(w http.ResponseWriter, r *http.Request) {
	sellers, err := usecases.ListSellerService.FindAll()
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
		map[string]interface{}{"sellers": sellers},
	)
}

func (s SellerController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	deleted, err := usecases.DeleteSellerService.Delete(int(id))

	if err != nil {
		WriteResponse(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	WriteResponse(
		w,
		http.StatusOK,
		map[string]interface{}{"deleted": deleted},
	)
}

var (
	Seller = SellerController{}
)
