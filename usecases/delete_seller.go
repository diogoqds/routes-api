package usecases

import (
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type DeleteSeller interface {
	Delete(id int) (bool, error)
}

type DeleteSellerUseCase struct {
}

func (d DeleteSellerUseCase) Delete(id int) (bool, error) {
	deleted, err := repositories.SellerRepo.DeleteSeller.Delete(id)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

var (
	DeleteSellerService DeleteSeller = DeleteSellerUseCase{}
)
