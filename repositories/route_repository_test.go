package repositories

import (
	"testing"
	"time"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoute_Success(t *testing.T) {
	setupTestDb()

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

	bounds := `{
		"type": "Polygon",
		"coordinates": [[
			[-104.05, 48.99],
			[-97.22,  48.98],
			[-96.58,  45.94],
			[-104.03, 45.94],
			[-104.05, 48.99]
		]]
	}`

	route, err := RouteRepo.RouteCreator.Create("route1", bounds, seller.Id)

	assert.EqualValues(t, 1, route.Id)
	assert.EqualValues(t, "route1", route.Name)
	assert.Nil(t, err)
}
