package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type UpdateRoute interface {
	Update(id int, name string, polygon entities.Polygon) (*entities.Route, error)
}

type UpdateRouteUseCase struct {
}

func (c UpdateRouteUseCase) Update(id int, name string, polygon entities.Polygon) (*entities.Route, error) {
	route, err := repositories.RouteRepo.RouteFinderById.FindById(id)

	if reflect.ValueOf(name).IsZero() {
		name = route.Name
	}

	polygonBytes, _ := json.Marshal(polygon)
	polygonString := fmt.Sprintf("%s", polygonBytes)

	if reflect.ValueOf(polygon).IsZero() {
		polygon = route.Bounds
		polygonBytes, _ = json.Marshal(route.Bounds)
		polygonString = fmt.Sprintf("%s", polygonBytes)
	}

	routes, err := repositories.RouteRepo.RouteFinder.FindByBounds(polygonString)

	if len(routes) > 0 {
		if routes[0].Id != route.Id {
			return nil, errors.New("There is already a route with these coordinates")
		}
	}

	route, err = repositories.RouteRepo.RouteUpdater.Update(id, name, polygonString)
	if err != nil {
		return nil, err
	}
	return route, nil
}

var (
	UpdateRouteService UpdateRoute = UpdateRouteUseCase{}
)
