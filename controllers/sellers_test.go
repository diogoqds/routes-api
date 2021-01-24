package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/middlewares"
	"github.com/diogoqds/routes-challenge-api/usecases"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockCreateSellerRepo struct {
	createFunc func(name string, email string) (*entities.Seller, error)
}

func (mock mockCreateSellerRepo) Create(name string, email string) (*entities.Seller, error) {
	return mock.createFunc(name, email)
}

func TestSellers_Create(t *testing.T) {
	scenarios := []struct {
		TestName   string
		Name       string
		Email      string
		StatusCode int
		createFunc func(name string, email string) (*entities.Seller, error)
	}{
		{
			TestName:   "when successfully creates a seller",
			Name:       "seller",
			Email:      "seller@email.com",
			StatusCode: http.StatusCreated,
			createFunc: func(name string, email string) (*entities.Seller, error) {
				return &entities.Seller{
					Id:        1,
					Name:      "seller",
					Email:     "seller@email.com",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: nil,
				}, nil
			},
		},
		{
			TestName:   "when name is empty",
			Name:       "",
			Email:      "seller@email.com",
			StatusCode: http.StatusUnprocessableEntity,
			createFunc: func(name string, email string) (*entities.Seller, error) {
				return nil, errors.New("name is required")
			},
		},
		{
			TestName:   "when email is empty",
			Name:       "seller",
			Email:      "",
			StatusCode: http.StatusUnprocessableEntity,
			createFunc: func(name string, email string) (*entities.Seller, error) {
				return nil, errors.New("email is required")
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			mockCreateSellerRepository := mockCreateSellerRepo{
				createFunc: scenario.createFunc,
			}

			usecases.CreateSellerService = mockCreateSellerRepository
			sellerController := SellerController{}
			params := fmt.Sprintf(`{ "name": "%s", "email": "%s" }`, scenario.Name, scenario.Email)
			requestBody := bytes.NewBuffer([]byte(params))

			request, _ := http.NewRequest("POST", "/sellers", requestBody)
			response := httptest.NewRecorder()

			setupAuth(request)

			middlewares.AuthMiddleware(sellerController.Create).ServeHTTP(response, request)
			assert.EqualValues(t, scenario.StatusCode, response.Code)
		})
	}
}
