package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockSellerRepository struct {
	deleteSellerFunc func(id int) (bool, error)
}

func (mock mockSellerRepository) Delete(id int) (bool, error) {
	return mock.deleteSellerFunc(id)
}

func TestDeleteSeller_Success(t *testing.T) {
	sellerRepository := mockSellerRepository{
		deleteSellerFunc: func(id int) (bool, error) {
			return true, nil
		},
	}

	repositories.SellerRepo.DeleteSeller = sellerRepository

	deleted, err := DeleteSellerService.Delete(1)

	assert.Nil(t, err)
	assert.EqualValues(t, true, deleted)
}

func TestDeleteSeller_Error(t *testing.T) {
	sellerRepository := mockSellerRepository{
		deleteSellerFunc: func(id int) (bool, error) {
			return false, errors.New("error while deleting seller")
		},
	}

	repositories.SellerRepo.DeleteSeller = sellerRepository

	deleted, err := DeleteSellerService.Delete(1)

	assert.EqualValues(t, "error while deleting seller", err.Error())
	assert.EqualValues(t, false, deleted)
}
