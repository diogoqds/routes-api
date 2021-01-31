package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/usecases"
)

type ClientController struct {
}

type RequestParams struct {
	Name        string         `json:"name"`
	Geolocation entities.Point `json:"geolocation"`
}

func (c ClientController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(
			w,
			http.StatusNotFound,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	var bodyParams RequestParams

	if err = json.Unmarshal(body, &bodyParams); err != nil {
		WriteResponse(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	client, err := usecases.CreateClientService.Create(bodyParams.Name, bodyParams.Geolocation)

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
		map[string]interface{}{"client": client},
	)

}

func (c ClientController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(
			w,
			http.StatusNotFound,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	var bodyParams RequestParams

	if err = json.Unmarshal(body, &bodyParams); err != nil {
		WriteResponse(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	client, err := usecases.UpdateClientService.Update(id, bodyParams.Name, bodyParams.Geolocation)

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
		map[string]interface{}{"client": client},
	)

}

func (c ClientController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	deleted, err := usecases.DeleteClientService.Delete(id)

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
	Client = ClientController{}
)
