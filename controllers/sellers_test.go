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

type mockCreateSellerUseCase struct {
	createFunc func(name string, email string) (*entities.Seller, error)
}

func (mock mockCreateSellerUseCase) Create(name string, email string) (*entities.Seller, error) {
	return mock.createFunc(name, email)
}

type mockListSellersUseCase struct {
	listFunc func() ([]entities.Seller, error)
}

func (mock mockListSellersUseCase) FindAll() ([]entities.Seller, error) {
	return mock.listFunc()
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
			mockCreateSellerService := mockCreateSellerUseCase{
				createFunc: scenario.createFunc,
			}

			usecases.CreateSellerService = mockCreateSellerService
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

func TestSellers_Index(t *testing.T) {
	scenarios := []struct {
		TestName   string
		listFunc   func() ([]entities.Seller, error)
		StatusCode int
	}{
		{
			TestName: "when successfully fetches the sellers",
			listFunc: func() ([]entities.Seller, error) {
				return []entities.Seller{
					{
						Id:        1,
						Name:      "Seller 1",
						Email:     "seller@email.com",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
						DeletedAt: nil,
					},
				}, nil
			},
			StatusCode: http.StatusOK,
		},
		{
			TestName: "when an error happens",
			listFunc: func() ([]entities.Seller, error) {
				return []entities.Seller{}, errors.New("error while fetching sellers")
			},
			StatusCode: http.StatusInternalServerError,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			mockListSellersService := mockListSellersUseCase{
				listFunc: scenario.listFunc,
			}

			usecases.ListSellerService = mockListSellersService
			sellerController := SellerController{}

			requestBody := bytes.NewBuffer([]byte(""))

			request, _ := http.NewRequest("GET", "/sellers", requestBody)
			response := httptest.NewRecorder()

			setupAuth(request)

			middlewares.AuthMiddleware(sellerController.Index).ServeHTTP(response, request)
			assert.EqualValues(t, scenario.StatusCode, response.Code)
		})
	}
}
