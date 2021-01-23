package repositories

import (
	"github.com/diogoqds/routes-challenge-api/entities"
)

type CreateSeller interface {
	Create(name string, email string) (*entities.Seller, error)
}

type SellerRepository struct {
	CreateSeller CreateSeller
}

type createSellerImplementation struct{}

func (c createSellerImplementation) Create(name string, email string) (*entities.Seller, error) {
	return nil, nil
}

var (
	SellerRepo = SellerRepository{
		CreateSeller: createSellerImplementation{},
	}
)
