package repositories

import (
	"encoding/json"
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

type RouteUpdater interface {
	Update(id int, name string, bounds string) (*entities.Route, error)
}

type RouteRepository struct {
	RouteCreator RouteCreator
	RouteFinder  RouteFinder
	RouteUpdater RouteUpdater
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

	if err != nil {
		log.Println("Error fetching the routes: " + err.Error())
		return nil, err
	}

	return routes, nil
}

func (r routeRepositoryImplementation) Update(id int, name string, bounds string) (*entities.Route, error) {
	var route entities.Route
	var polygon entities.Polygon

	query := "UPDATE routes SET name = $1, bounds = ST_GeomFromGeoJSON($2::text) WHERE id = $3 RETURNING id, name, created_at, updated_at, deleted_at;"

	err := infra.DB.QueryRow(query, name, bounds, id).
		Scan(&route.Id, &route.Name, &route.CreatedAt, &route.UpdatedAt, &route.DeletedAt)

	err = json.Unmarshal([]byte(bounds), &polygon)

	if err != nil {
		log.Println("Error unmarshalling the bounds: " + err.Error())
		return nil, err
	}

	route.Id = id
	route.Bounds = &polygon

	if err != nil {
		log.Println("Error updating the route: " + err.Error())
		return nil, err
	}

	return &route, nil
}

var (
	RouteRepo = RouteRepository{
		RouteCreator: routeRepositoryImplementation{},
		RouteFinder:  routeRepositoryImplementation{},
		RouteUpdater: routeRepositoryImplementation{},
	}
)
