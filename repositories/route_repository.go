package repositories

import (
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	g "github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
	"log"
	"time"
)

type RouteCreator interface {
	Create(name string, bounds *g.Bounds, sellerId int) (*entities.Route, error)
}

type RouteRepository struct {
	RouteCreator RouteCreator
}

type routeRepositoryImplementation struct{}

func (r routeRepositoryImplementation) Create(name string, bounds *g.Bounds, sellerId int) (*entities.Route, error) {
	route := entities.Route{
		Name:      name,
		SellerId:  sellerId,
		Bounds:    bounds,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
	sqlInsert := "INSERT INTO routes (name, bounds, seller_id) VALUES ($1, $2, $3) RETURNING id"
	var id int

	bytes, err := geojson.Marshal(bounds.Polygon())
	if err != nil {
		log.Println("Error marshalling the bounds: " + err.Error())
		return nil, err
	}

	err = infra.DB.QueryRow(sqlInsert, name, bytes, sellerId).Scan(&id)

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
