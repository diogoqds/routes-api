package usecases

import (
	"database/sql"

	"github.com/diogoqds/routes-challenge-api/repositories"
	"github.com/diogoqds/routes-challenge-api/validators"
)

type AssociateSeller interface {
	Associate(routeId int, sellerId int) (bool, error)
}

type AssociateSellerUseCase struct {
}

func (d AssociateSellerUseCase) Associate(routeId int, sellerId int) (bool, error) {
	err := validators.RouteSellerValidator.RouteWithSellerId(sellerId)
	if err != sql.ErrNoRows {
		return false, err
	}
	return repositories.RouteRepo.RouteSellerUpdater.Associate(routeId, sellerId)
}

var (
	AssociateSellerService AssociateSeller = AssociateSellerUseCase{}
)
