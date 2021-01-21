package controllers

import (
	"bytes"
	"errors"
	"github.com/diogoqds/routes-challenge-api/usecases"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Scenario struct {
	TestName         string
	authenticateFunc func(email string) (string, error)
	StatusCode       int
}

type mockAuth struct {
	authenticateFunc func(email string) (string, error)
}

func (mock mockAuth) Authenticate(email string) (string, error) {
	return mock.authenticateFunc(email)
}

func TestLogin(t *testing.T) {
	scenarios := []Scenario{
		{
			TestName: "when successfully generate a token",
			authenticateFunc: func(email string) (string, error) {

				return "token", nil
			},
			StatusCode: http.StatusCreated,
		},
		{
			TestName: "when failing to generate a token",
			authenticateFunc: func(email string) (string, error) {
				return "", errors.New("error while generate the token")
			},
			StatusCode: http.StatusBadRequest,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			mockAuthService := mockAuth{}
			mockAuthService.authenticateFunc = scenario.authenticateFunc
			usecases.AuthService = mockAuthService

			requestBody := bytes.NewBuffer([]byte(`{"email":"admin@email.com"}`))

			request, _ := http.NewRequest("POST", "/login", requestBody)
			response := httptest.NewRecorder()

			Login(response, request)

			assert.EqualValues(t, scenario.StatusCode, response.Code)
		})
	}
}
