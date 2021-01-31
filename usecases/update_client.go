package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/diogoqds/routes-challenge-api/repositories"

	"github.com/diogoqds/routes-challenge-api/entities"
)

type UpdateClient interface {
	Update(id int, name string, point entities.Point) (*entities.Client, error)
}

type UpdateClientUseCase struct {
}

func (c UpdateClientUseCase) Update(id int, name string, point entities.Point) (*entities.Client, error) {
	client, err := repositories.ClientRepo.ClientFinderById.FindById(id)

	if err != nil {
		return nil, err
	}
	pointBytes, _ := json.Marshal(point)
	geolocationString := fmt.Sprintf("%s", pointBytes)

	if name == "" {
		name = client.Name
	}

	if reflect.ValueOf(point).IsZero() {
		point = client.Geolocation
		pointBytes, _ = json.Marshal(client.Geolocation)
		geolocationString = fmt.Sprintf("%s", pointBytes)
	}

	route, err := FindClientRouteService.Find(point)

	if err != nil {
		log.Println("error fetching the client route", err.Error())
		return nil, err
	}

	client, err = repositories.ClientRepo.ClientUpdater.Update(id, name, geolocationString, route.Id)

	if err != nil {
		log.Println("error updating the client", err.Error())
		return nil, err
	}

	return client, nil
}

var (
	UpdateClientService UpdateClient = UpdateClientUseCase{}
)
