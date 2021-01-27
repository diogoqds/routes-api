package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
	g "github.com/twpayne/go-geom"
)

type CreateRoute interface {
	Create(name string, bounds *g.Bounds, sellerId int) (*entities.Route, error)
}

type CreateRouteUseCase struct {
}

func (c CreateRouteUseCase) Create(name string, bounds *g.Bounds, sellerId int) (*entities.Route, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if sellerId == 0 {
		return nil, errors.New("seller_id is required")
	}

	if bounds == nil {
		return nil, errors.New("bounds is required")
	}

	route, err := repositories.RouteRepo.RouteCreator.Create(name, bounds, sellerId)
	if err != nil {
		return nil, err
	}

	return route, nil
}

var (
	CreateRouteService CreateRoute = CreateRouteUseCase{}
)
