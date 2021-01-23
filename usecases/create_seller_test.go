package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockSellerRepo struct {
	createSellerFunc func(name string, email string) (*entities.Seller, error)
}

func (mock mockSellerRepo) Create(name string, email string) (*entities.Seller, error) {
	return mock.createSellerFunc(name, email)
}

func TestCreateSeller_Failure(t *testing.T) {
	scenarios := []struct {
		TestName         string
		Name             string
		Email            string
		CreateSellerFunc func(name string, email string) (*entities.Seller, error)
		Seller           *entities.Seller
		ErrorMessage     string
	}{
		{
			TestName:         "when name is empty",
			Name:             "",
			Email:            "seller@email.com",
			CreateSellerFunc: nil,
			Seller:           nil,
			ErrorMessage:     "name is required",
		},
		{
			TestName:         "when email is empty",
			Name:             "valid_name",
			Email:            "",
			CreateSellerFunc: nil,
			Seller:           nil,
			ErrorMessage:     "email is required",
		},
		{
			TestName: "when an error occurs in the persistence",
			Name:     "valid_name",
			Email:    "valid_email",
			CreateSellerFunc: func(name string, email string) (*entities.Seller, error) {
				return nil, errors.New("error while saving the seller")
			},
			Seller:       nil,
			ErrorMessage: "error while saving the seller",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			sellerRepository := mockSellerRepo{
				createSellerFunc: scenario.CreateSellerFunc,
			}

			repositories.SellerRepo.CreateSeller = sellerRepository

			seller, err := CreateSellerService.Create(scenario.Name, scenario.Email)

			assert.Nil(t, seller)
			assert.EqualValues(t, scenario.ErrorMessage, err.Error())
		})
	}
}

func TestCreateSeller_Success(t *testing.T) {
	sellerRepository := mockSellerRepo{
		createSellerFunc: func(name string, email string) (*entities.Seller, error) {
			return &entities.Seller{
				Id:        1,
				Name:      "seller",
				Email:     "seller@email.com",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			}, nil
		},
	}

	repositories.SellerRepo.CreateSeller = sellerRepository

	seller, err := CreateSellerService.Create("seller", "seller@email.com")

	assert.Nil(t, err)
	assert.EqualValues(t, "seller", seller.Name)
	assert.EqualValues(t, "seller@email.com", seller.Email)
	assert.NotNil(t, seller.Id)
	assert.NotNil(t, seller.CreatedAt)
	assert.NotNil(t, seller.UpdatedAt)
	assert.Nil(t, seller.DeletedAt)
}
