package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	"log"
	"time"
)

type ClientCreator interface {
	Create(name string, geolocation string, routeId int) (*entities.Client, error)
}

type ClientUpdater interface {
	Update(id int, name string, geolocation string, routeId int) (*entities.Client, error)
}
type ClientFinderById interface {
	FindById(id int) (*entities.Client, error)
}

type ClientRepository struct {
	ClientCreator    ClientCreator
	ClientUpdater    ClientUpdater
	ClientFinderById ClientFinderById
}

type clientRepository struct{}

func (c clientRepository) FindById(id int) (*entities.Client, error) {
	var client entities.Client
	var geolocationString string

	query := "SELECT id, name, ST_AsGeoJSON(geolocation) as geolocation, created_at, updated_at, deleted_at FROM clients WHERE id = $1"
	err := infra.DB.QueryRow(query, id).Scan(&client.Id, &client.Name, &geolocationString, &client.CreatedAt, &client.UpdatedAt, &client.DeletedAt)

	json.Unmarshal([]byte(geolocationString), &client.Geolocation)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no client with id %d\n", id)
		return nil, err
	case err != nil:
		log.Printf("query error: %v\n", err)
		return nil, err
	default:
		return &client, nil
	}
}

func (c clientRepository) Create(name string, geolocation string, routeId int) (*entities.Client, error) {

	client := entities.Client{
		Name:      name,
		RouteId:   routeId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	sql := "INSERT INTO clients (name, geolocation, route_id) VALUES ($1, ST_GeomFromGeoJSON($2::text), $3) RETURNING id"
	var id int

	err := infra.DB.QueryRow(sql, name, geolocation, routeId).Scan(&id)

	if err != nil {
		log.Println("Error saving the client: " + err.Error())
		return nil, err
	}

	client.Id = id

	return &client, nil
}

func (c clientRepository) Update(id int, name string, geolocation string, routeId int) (*entities.Client, error) {
	var client entities.Client
	sql := "UPDATE clients SET name = $1, geolocation = ST_GeomFromGeoJSON($2::text), route_id = $3 WHERE id = $4 RETURNING id"

	err := infra.DB.QueryRow(sql, name, geolocation, routeId, id).
		Scan(&client.Id)

	if err != nil {
		log.Println("error while updating client", err.Error())
		return nil, errors.New("error while updating client")
	}

	return &client, nil
}

var (
	ClientRepo = ClientRepository{
		ClientCreator:    clientRepository{},
		ClientUpdater:    clientRepository{},
		ClientFinderById: clientRepository{},
	}
)
