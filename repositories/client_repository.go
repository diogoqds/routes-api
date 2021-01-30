package repositories

import (
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	"log"
	"time"
)

type ClientCreator interface {
	Create(name string, geolocation string, routeId int) (*entities.Client, error)
}

type ClientRepository struct {
	ClientCreator ClientCreator
}

type clientRepository struct{}

func (c clientRepository) Create(name string, geolocation string, routeId int) (*entities.Client, error) {

	client := entities.Client{
		Name:      name,
		RouteId:   routeId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	sql := "INSERT INTO clients (name, geolocation, route_id) VALUES ($1, ST_GeomFromGeoJSON($2::text), $3)"
	var id int

	err := infra.DB.QueryRow(sql, name, geolocation, routeId).Scan(&id)

	if err != nil {
		log.Println("Error saving the client: " + err.Error())
		return nil, err
	}

	client.Id = id

	return &client, nil
}

var (
	ClientRepo = ClientRepository{
		ClientCreator: clientRepository{},
	}
)
