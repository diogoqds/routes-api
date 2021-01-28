package repositories

import (
	"log"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"

	"time"
)

type RouteCreator interface {
	Create(name string, bounds string, sellerId int) (*entities.Route, error)
}

type RouteRepository struct {
	RouteCreator RouteCreator
}

type routeRepositoryImplementation struct{}

func (r routeRepositoryImplementation) Create(name string, bounds string, sellerId int) (*entities.Route, error) {
	route := entities.Route{
		Name:      name,
		SellerId:  sellerId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	sqlInsert := `INSERT INTO routes (name, bounds, seller_id) VALUES ($1, ST_GeomFromGeoJSON($2::text), $3) RETURNING id`
	var id int

	err := infra.DB.QueryRow(sqlInsert, name, bounds, sellerId).Scan(&id)

	if err != nil {
		log.Println("Error saving the route: " + err.Error())
		return nil, err
	}

	route.Id = id
	return &route, nil
}

var (
	RouteRepo = RouteRepository{
		RouteCreator: routeRepositoryImplementation{},
	}
)
