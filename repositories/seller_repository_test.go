package repositories

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	query = "INSERT INTO sellers"
)

func TestCreateSeller_Success(t *testing.T) {
	setupDb()

	rows := mock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(query).
		WithArgs("seller", "seller@email.com").
		WillReturnRows(rows)

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, 1, seller.Id)
	assert.EqualValues(t, "seller@email.com", seller.Email)
	assert.Nil(t, err)
}

func TestCreateSeller_ErrorSaving(t *testing.T) {
	setupDb()
	mock.ExpectQuery(query).
		WithArgs("seller", "seller@email.com").
		WillReturnError(errors.New("generic error"))

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, "generic error", err.Error())
	assert.Nil(t, seller)
}

func TestCreateSeller_ErrorReturningId(t *testing.T) {
	setupDb()
	mock.ExpectQuery(query).
		WithArgs("seller", "seller@email.com").
		WillReturnError(errors.New("result error"))

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, "result error", err.Error())
	assert.Nil(t, seller)
}
