package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type CreateSeller interface {
	Create(name string, email string) (*entities.Seller, error)
}

type CreateSellerUseCase struct {
}

func (c CreateSellerUseCase) Create(name string, email string) (*entities.Seller, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}

	seller, err := repositories.SellerRepo.CreateSeller.Create(name, email)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

var (
	CreateSellerService CreateSeller = CreateSellerUseCase{}
)
