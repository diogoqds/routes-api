package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/diogoqds/routes-challenge-api/usecases"
)

type requestParams struct {
	Email string `json:"email"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(
			w,
			http.StatusNotFound,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	var bodyParams requestParams
	if err = json.Unmarshal(body, &bodyParams); err != nil {
		WriteResponse(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"message": err.Error()},
		)
		return
	}

	token, err := usecases.AuthService.Authenticate(bodyParams.Email)
	if err != nil {
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
	}

	WriteResponse(
		w,
		http.StatusCreated,
		map[string]interface{}{"token": token},
	)
	w.WriteHeader(http.StatusCreated)
}
