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

type RoutesController struct {
}

type RouteRequestParams struct {
	Name     string           `json:"name"`
	Bounds   entities.Polygon `json:"bounds"`
	SellerId int              `json:"seller_id"`
}

func (c RoutesController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(
			w,
			http.StatusNotFound,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	var bodyParams RouteRequestParams

	if err = json.Unmarshal(body, &bodyParams); err != nil {
		WriteResponse(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	if err != nil {
		WriteResponse(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	route, err := usecases.CreateRouteService.Create(bodyParams.Name, bodyParams.Bounds, bodyParams.SellerId)

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
		map[string]interface{}{"route": route},
	)

}

func (c RoutesController) Update(w http.ResponseWriter, r *http.Request) {
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

	var bodyParams RouteRequestParams

	if err = json.Unmarshal(body, &bodyParams); err != nil {
		WriteResponse(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	route, err := usecases.UpdateRouteService.Update(id, bodyParams.Name, bodyParams.Bounds)

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
		http.StatusOK,
		map[string]interface{}{"route": route},
	)
}

func (c RoutesController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	deleted, err := usecases.DeleteRouteService.Delete(id)

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

func (c RoutesController) AssociateSeller(w http.ResponseWriter, r *http.Request) {
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

	var bodyParams RouteRequestParams

	if err = json.Unmarshal(body, &bodyParams); err != nil {
		WriteResponse(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	if err != nil {
		WriteResponse(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	associated, err := usecases.AssociateSellerService.Associate(id, bodyParams.SellerId)

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
		map[string]interface{}{"associated": associated},
	)
}

func (c RoutesController) DisassociateSeller(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	disassociated, err := usecases.DisassociateSellerService.Disassociate(id)

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
		map[string]interface{}{"disassociated": disassociated},
	)
}

var (
	Routes = RoutesController{}
)
