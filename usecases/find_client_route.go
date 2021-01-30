package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"log"
)

type FindClientRoute interface {
	Find(point entities.Point) (*entities.Route, error)
}

type FindClientRouteUseCase struct {
}

func (f FindClientRouteUseCase) Find(point entities.Point) (*entities.Route, error) {
	const defaultRouteName = "Outros"

	pointBytes, _ := json.Marshal(point)
	pointString := fmt.Sprintf("%s", pointBytes)

	route, err := repositories.RouteRepo.RouteFinderByPoint.FindByPoint(pointString)

	if err != nil {
		log.Println("usecase error, fetching route by point", err.Error())
		return nil, err
	}

	if route == nil {
		route, err = repositories.RouteRepo.RouteFinderByName.FindByName(defaultRouteName)
		if err != nil {
			log.Println("usecase error, fetching route by name", err.Error())
			return nil, err
		}
	}

	return route, nil
}

var (
	FindClientRouteService FindClientRoute = FindClientRouteUseCase{}
)
