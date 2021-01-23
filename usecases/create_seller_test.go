package usecases

import (
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSeller_Failure(t *testing.T) {
	scenarios := []struct {
		TestName         string
		Name             string
		Email            string
		CreateSellerFunc func(name string, email string) (*entities.Seller, error)
		Seller           *entities.Seller
		ErrorMessage     string
	}{
		{
			TestName:         "when name is empty",
			Name:             "",
			Email:            "seller@email.com",
			CreateSellerFunc: nil,
			Seller:           nil,
			ErrorMessage:     "name is required",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			seller, err := CreateSellerService.Create(scenario.Name, scenario.Email)

			assert.Nil(t, seller)
			assert.EqualValues(t, scenario.ErrorMessage, err.Error())
		})
	}
}
