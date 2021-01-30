package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"log"

	"github.com/diogoqds/routes-challenge-api/entities"
)

type CreateClient interface {
	Create(name string, point entities.Point) (*entities.Client, error)
}

type CreateClientUseCase struct {
}

func (c CreateClientUseCase) Create(name string, point entities.Point) (*entities.Client, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	pointBytes, _ := json.Marshal(point)
	geolocationString := fmt.Sprintf("%s", pointBytes)

	route, err := FindClientRouteService.Find(point)

	if err != nil {
		log.Println("error fetching the client route", err.Error())
		return nil, err
	}

	client, err := repositories.ClientRepo.ClientCreator.Create(name, geolocationString, route.Id)

	if err != nil {
		log.Println("error creating the client", err.Error())
		return nil, err
	}

	return client, nil
}

var (
	CreateClientService CreateClient = CreateClientUseCase{}
)
