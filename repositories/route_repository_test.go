package repositories

import (
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/stretchr/testify/assert"
	g "github.com/twpayne/go-geom"
	"testing"
	"time"
)

func TestCreateRoute_Success(t *testing.T) {
	setupTestDb()
	bounds := g.Bounds{}
	seller := entities.Seller{
		Id:        1,
		Name:      "seller",
		Email:     "seller@email",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	query := "INSERT INTO routes"
	rows := mock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(query).
		WillReturnRows(rows)

	route, err := RouteRepo.RouteCreator.Create("route1", &bounds, &seller)

	assert.EqualValues(t, 1, route.Id)
	assert.EqualValues(t, "route1", route.Name)
	assert.Nil(t, err)
}
