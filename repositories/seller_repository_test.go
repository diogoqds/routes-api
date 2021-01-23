package repositories

import (
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestCreateSeller_Success(t *testing.T) {
	var db *sqlx.DB
	var err error
	var mock sqlmock.Sqlmock

	db, mock, err = sqlmock.Newx()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}

	infra.DB = db

	query := "INSERT INTO sellers"

	mock.ExpectExec(query).
		WithArgs("seller", "seller@email.com").
		WillReturnResult(sqlmock.NewResult(1, 1))

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, 1, seller.Id)
	assert.EqualValues(t, "seller@email.com", seller.Email)
	assert.Nil(t, err)
}
