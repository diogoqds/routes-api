package repositories

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

var (
	db    *sqlx.DB
	err   error
	mock  sqlmock.Sqlmock
	query = "INSERT INTO sellers"
)

func setupDb() {
	db, mock, err = sqlmock.Newx()

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}

	infra.DB = db
}

func TestCreateSeller_Success(t *testing.T) {
	setupDb()
	mock.ExpectExec(query).
		WithArgs("seller", "seller@email.com").
		WillReturnResult(sqlmock.NewResult(1, 1))

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, 1, seller.Id)
	assert.EqualValues(t, "seller@email.com", seller.Email)
	assert.Nil(t, err)
}

func TestCreateSeller_ErrorSaving(t *testing.T) {
	setupDb()
	mock.ExpectExec(query).
		WithArgs("seller", "seller@email.com").
		WillReturnError(errors.New("generic error"))

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, "generic error", err.Error())
	assert.Nil(t, seller)
}

func TestCreateSeller_ErrorReturningId(t *testing.T) {
	setupDb()
	mock.ExpectExec(query).
		WithArgs("seller", "seller@email.com").
		WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, "result error", err.Error())
	assert.Nil(t, seller)
}
