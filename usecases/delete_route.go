package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type DeleteRoute interface {
	Delete(id int) (bool, error)
}

type DeleteRouteUseCase struct {
}

func (d DeleteRouteUseCase) Delete(id int) (bool, error) {
	route, err := repositories.RouteRepo.RouteFinderById.FindById(id)

	if err != nil {
		return false, err
	}

	if route.SellerId != 0 {
		return false, errors.New("cannot delete while route has a seller")
	}

	deleted, err := repositories.RouteRepo.RouteEraser.Delete(id)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

var (
	DeleteRouteService DeleteRoute = DeleteRouteUseCase{}
)
