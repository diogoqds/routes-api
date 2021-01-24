package repositories

import (
	"errors"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"regexp"
	"testing"
	"time"
)

func TestCreateSeller_Success(t *testing.T) {
	setupTestDb()

	query := "INSERT INTO sellers"
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
	setupTestDb()
	query := "INSERT INTO sellers"
	mock.ExpectQuery(query).
		WithArgs("seller", "seller@email.com").
		WillReturnError(errors.New("generic error"))

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, "generic error", err.Error())
	assert.Nil(t, seller)
}

func TestCreateSeller_ErrorReturningId(t *testing.T) {
	setupTestDb()
	query := "INSERT INTO sellers"
	mock.ExpectQuery(query).
		WithArgs("seller", "seller@email.com").
		WillReturnError(errors.New("result error"))

	seller, err := SellerRepo.CreateSeller.Create("seller", "seller@email.com")

	assert.EqualValues(t, "result error", err.Error())
	assert.Nil(t, seller)
}

func TestListSeller_Success(t *testing.T) {
	scenarios := []struct {
		TestName string
		Quantity int
		MockData func() *sqlmock.Rows
	}{
		{
			TestName: "when return sellers",
			Quantity: 2,
			MockData: func() *sqlmock.Rows {
				rows := mock.NewRows(
					[]string{"id", "name", "email", "created_at", "updated_at", "deleted_at"},
				).
					AddRow(1, "Seller 1", "seller1@email.com", time.Time{}, time.Time{}, nil).
					AddRow(2, "Seller 2", "seller2@email.com", time.Time{}, time.Time{}, nil)

				mock.NewRows(
					[]string{"id", "name", "email", "created_at", "updated_at", "deleted_at"},
				).AddRow(3, "Seller 3", "seller3@email.com", time.Time{}, time.Time{}, time.Time{})

				return rows
			},
		},
		{
			TestName: "when there aren't sellers",
			Quantity: 0,
			MockData: func() *sqlmock.Rows {
				return &sqlmock.Rows{}
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			setupTestDb()

			rows := scenario.MockData()
			query := "SELECT * FROM sellers WHERE deleted_at IS NULL"
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			sellers, err := SellerRepo.ListSellers.FindAll()

			assert.Nil(t, err)
			assert.Equal(t, scenario.Quantity, len(sellers))
		})
	}
}

func TestListSeller_Error(t *testing.T) {

	setupTestDb()

	query := "SELECT * FROM sellers WHERE deleted_at IS NULL"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.New("error while fetching sellers"))

	sellers, err := SellerRepo.ListSellers.FindAll()

	assert.EqualValues(t, "error while fetching sellers", err.Error())
	assert.Nil(t, sellers)

}

func TestDeleteSeller_Success(t *testing.T) {
	setupTestDb()
	rows := mock.NewRows([]string{"id"}).AddRow(1)

	query := "UPDATE sellers SET deleted_at = NOW() WHERE id = $1 RETURNING id"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	deleted, err := SellerRepo.DeleteSeller.Delete(1)
	assert.Nil(t, err)
	assert.EqualValues(t, true, deleted)
}

func TestDeleteSeller_Error(t *testing.T) {
	setupTestDb()

	query := "UPDATE sellers SET deleted_at = NOW() WHERE id = $1 RETURNING id"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error deleting seller"))

	deleted, err := SellerRepo.DeleteSeller.Delete(1)
	assert.EqualValues(t, "error deleting seller", err.Error())
	assert.EqualValues(t, false, deleted)

}
