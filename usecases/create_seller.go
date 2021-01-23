package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/entities"
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
	return nil, nil
}

var (
	CreateSellerService CreateSeller = CreateSellerUseCase{}
)
