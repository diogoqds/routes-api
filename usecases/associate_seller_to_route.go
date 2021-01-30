package usecases

import (
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type AssociateSeller interface {
	Associate(routeId int, sellerId int) (bool, error)
}

type AssociateSellerUseCase struct {
}

func (d AssociateSellerUseCase) Associate(routeId int, sellerId int) (bool, error) {
	return repositories.RouteRepo.RouteSellerUpdater.Associate(routeId, sellerId)
}

var (
	AssociateSellerService AssociateSeller = AssociateSellerUseCase{}
)
