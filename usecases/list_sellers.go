package usecases

import (
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type ListSellers interface {
	FindAll() ([]entities.Seller, error)
}

type ListSellersUseCase struct {
}

func (l ListSellersUseCase) FindAll() ([]entities.Seller, error) {
	sellers, err := repositories.SellerRepo.ListSellers.FindAll()

	if err != nil {
		return []entities.Seller{}, err
	}

	return sellers, nil
}

var (
	ListSellerService ListSellers = ListSellersUseCase{}
)
