package usecases

import (
	"errors"
	"fmt"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type CreateRoute interface {
	Create(name string, bounds interface{}, sellerId int) (*entities.Route, error)
}

type CreateRouteUseCase struct {
}

func (c CreateRouteUseCase) Create(name string, bounds interface{}, sellerId int) (*entities.Route, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if sellerId == 0 {
		return nil, errors.New("seller_id is required")
	}
	boundsString := fmt.Sprintf("%v", bounds)

	route, err := repositories.RouteRepo.RouteCreator.Create(name, boundsString, sellerId)
	if err != nil {
		return nil, err
	}

	return route, nil
}

var (
	CreateRouteService CreateRoute = CreateRouteUseCase{}
)
