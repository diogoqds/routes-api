package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/usecases"
)

type RoutesController struct {
}

type BoundsParam struct {
	Bounds entities.Polygon `json:"bounds"`
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

	sellerId, err := strconv.ParseInt(fmt.Sprintf("%.0f", bodyParams["seller_id"]), 10, 64)

	if err != nil {
		WriteResponse(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	var b BoundsParam
	err = json.Unmarshal(body, &b)

	if err != nil {
		WriteResponse(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	route, err := usecases.CreateRouteService.Create(name, b.Bounds, int(sellerId))

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

var (
	Routes = RoutesController{}
)
