package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type UpdateRoute interface {
	Update(id int, name string, polygon entities.Polygon) (*entities.Route, error)
}

type UpdateRouteUseCase struct {
}

func (c UpdateRouteUseCase) Update(id int, name string, polygon entities.Polygon) (*entities.Route, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}

	if name == "" {
		return nil, errors.New("name is required")
	}

	polygonBytes, _ := json.Marshal(polygon)
	polygonString := fmt.Sprintf("%s", polygonBytes)

	routes, err := repositories.RouteRepo.RouteFinder.FindByBounds(polygonString)

	if len(routes) > 0 {
		return nil, errors.New("There is already a route with these coordinates")
	}

	route, err := repositories.RouteRepo.RouteUpdater.Update(id, name, polygonString)
	if err != nil {
		return nil, err
	}
	return route, nil
}

var (
	UpdateRouteService UpdateRoute = UpdateRouteUseCase{}
)
