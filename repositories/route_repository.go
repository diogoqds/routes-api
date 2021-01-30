package repositories

import (
	"fmt"
	"log"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"

	"time"
)

type RouteCreator interface {
	Create(name string, bounds string, sellerId int) (*entities.Route, error)
}

type RouteFinder interface {
	FindByBounds(bounds string) ([]entities.Route, error)
}

type RouteRepository struct {
	RouteCreator RouteCreator
	RouteFinder  RouteFinder
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

func (r routeRepositoryImplementation) FindByBounds(bounds string) ([]entities.Route, error) {
	routes := make([]entities.Route, 0)

	sql := `SELECT id FROM routes WHERE ST_INTERSECTS(ST_GeomFromGeoJson($1), routes.bounds)`

	err := infra.DB.Select(&routes, sql, bounds)

	fmt.Println("bounds", bounds)
	if err != nil {
		log.Println("Error fetching the routes: " + err.Error())
		return nil, err
	}

	return routes, nil
}

var (
	RouteRepo = RouteRepository{
		RouteCreator: routeRepositoryImplementation{},
		RouteFinder:  routeRepositoryImplementation{},
	}
)
