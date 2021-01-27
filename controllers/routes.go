package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/diogoqds/routes-challenge-api/usecases"
	g "github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RoutesController struct {
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
	boundsParam := fmt.Sprintf("%s", bodyParams["bounds"])

	var bounds g.T

	err = geojson.Unmarshal([]byte(boundsParam), &bounds)

	if err != nil {
		WriteResponse(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	sellerId, err := strconv.ParseInt(fmt.Sprintf("%.0f", bodyParams["seller_id"]), 10, 64)

	route, err := usecases.CreateRouteService.Create(name, bounds.Bounds(), int(sellerId))

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
