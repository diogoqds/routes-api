package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockListSellerRepository struct {
	findAllFunc func() ([]entities.Seller, error)
}

func (mock mockListSellerRepository) FindAll() ([]entities.Seller, error) {
	return mock.findAllFunc()
}

func TestListSellers_Success(t *testing.T) {
	scenarios := []struct {
		TestName    string
		findAllFunc func() ([]entities.Seller, error)
		Quantity    int
	}{
		{
			TestName: "when there are sellers",
			findAllFunc: func() ([]entities.Seller, error) {
				return []entities.Seller{
					{
						Id:        1,
						Name:      "Seller 1",
						Email:     "seller1@email.com",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
						DeletedAt: nil,
					},
					{
						Id:        2,
						Name:      "Seller 2",
						Email:     "seller2@email.com",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
						DeletedAt: nil,
					},
				}, nil
			},
			Quantity: 2,
		},
		{
			TestName: "when there are not sellers",
			findAllFunc: func() ([]entities.Seller, error) {
				return []entities.Seller{}, nil
			},
			Quantity: 0,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {

			mockListSellerService := mockListSellerRepository{}
			mockListSellerService.findAllFunc = scenario.findAllFunc

			repositories.SellerRepo.ListSellers = mockListSellerService

			sellers, err := ListSellerService.FindAll()
			assert.Nil(t, err)
			assert.EqualValues(t, scenario.Quantity, len(sellers))
		})
	}
}

func TestListSellers_Error(t *testing.T) {

	mockListSellerService := mockListSellerRepository{}
	mockListSellerService.findAllFunc = func() ([]entities.Seller, error) {
		return []entities.Seller{}, errors.New("error while fetching sellers")
	}

	repositories.SellerRepo.ListSellers = mockListSellerService

	sellers, err := ListSellerService.FindAll()
	assert.EqualValues(t, "error while fetching sellers", err.Error())
	assert.EqualValues(t, 0, len(sellers))

}
