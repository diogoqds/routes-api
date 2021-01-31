package usecases

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/diogoqds/routes-challenge-api/validators"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type CreateRoute interface {
	Create(name string, polygon entities.Polygon, sellerId int) (*entities.Route, error)
}

type CreateRouteUseCase struct {
}

func (c CreateRouteUseCase) Create(name string, polygon entities.Polygon, sellerId int) (*entities.Route, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if sellerId == 0 {
		return nil, errors.New("seller_id is required")
	}

	err := validators.RouteSellerValidator.RouteWithSellerId(sellerId)

	if err != sql.ErrNoRows {
		return nil, err
	}

	polygonBytes, _ := json.Marshal(polygon)
	polygonString := fmt.Sprintf("%s", polygonBytes)

	routes, err := repositories.RouteRepo.RouteFinder.FindByBounds(polygonString)

	if len(routes) > 0 {
		return nil, errors.New("There is already a route with these coordinates")
	}

	route, err := repositories.RouteRepo.RouteCreator.Create(name, polygonString, sellerId)
	if err != nil {
		return nil, err
	}

	return route, nil
}

var (
	CreateRouteService CreateRoute = CreateRouteUseCase{}
)
