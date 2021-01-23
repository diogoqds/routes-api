package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Scenario struct {
	Name          string
	AuthHeader    string
	JwtDecodeFunc func(token string) (jwt.MapClaims, error)
	FindByIdFunc  func(id int64) (*entities.Admin, error)
	StatusCode    int
	ErrorMessage  string
}

type mockJwtDecoder struct {
	decodeFunc func(token string) (jwt.MapClaims, error)
}

type mockAdminRepo struct {
	findByIdFunc func(id int64) (*entities.Admin, error)
}

func (mock mockJwtDecoder) Decode(token string) (jwt.MapClaims, error) {
	return mock.decodeFunc(token)
}

func (mock mockAdminRepo) FindById(id int64) (*entities.Admin, error) {
	return mock.findByIdFunc(id)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"hello": "world"}); err != nil {
		log.Println("[Response Body Error]", err)
		return
	}
}

func TestAuthMiddleware(t *testing.T) {
	scenarios := []Scenario{
		{
			Name:          "when token is in the wrong format",
			AuthHeader:    "",
			JwtDecodeFunc: nil,
			FindByIdFunc:  nil,
			StatusCode:    http.StatusUnauthorized,
			ErrorMessage:  "Malformed Token",
		},
		{
			Name:       "when jwt decoder returns an error",
			AuthHeader: "Bearer token",
			JwtDecodeFunc: func(token string) (jwt.MapClaims, error) {
				return nil, errors.New("error while decoding the token")
			},
			FindByIdFunc: nil,
			StatusCode:   http.StatusUnauthorized,
			ErrorMessage: "unauthorized",
		},
		{
			Name:       "when AdminRepository returns an error",
			AuthHeader: "Bearer token",
			JwtDecodeFunc: func(token string) (jwt.MapClaims, error) {
				return jwt.MapClaims{"id": float64(1)}, nil
			},
			FindByIdFunc: func(id int64) (*entities.Admin, error) {
				return nil, errors.New("error while fetching the admin")
			},
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "error while fetching the admin",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			mockJwt := mockJwtDecoder{decodeFunc: scenario.JwtDecodeFunc}
			mockAdminRepository := mockAdminRepo{findByIdFunc: scenario.FindByIdFunc}

			infra.Jwt.Decoder = mockJwt
			repositories.AdminRepo.FinderById = mockAdminRepository

			request, _ := http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte("")))

			request.Header.Set("Authorization", scenario.AuthHeader)

			response := httptest.NewRecorder()

			AuthMiddleware(Hello).ServeHTTP(response, request)
			respBody, _ := ioutil.ReadAll(response.Body)

			var body map[string]interface{}

			json.Unmarshal(respBody, &body)
			fmt.Println("body>>>>>", body)
			assert.EqualValues(t, scenario.StatusCode, response.Code)
			assert.EqualValues(t, scenario.ErrorMessage, body["message"])
		})
	}

}
