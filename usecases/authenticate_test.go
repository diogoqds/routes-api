package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/diogoqds/routes-challenge-api/entities"
	"time"
)

type Scenario struct {
	TestName string
	Email string
	Token string
	Err error
	findByEmailFunc func(email string) (*entities.Admin, error)
}

type mockAdminRepo struct {
	findByEmailFunc func(email string) (*entities.Admin, error)
}

func (mock mockAdminRepo) FindByEmail(email string) (*entities.Admin, error) {
	return mock.findByEmailFunc(email)
}

func TestAuthenticate(t *testing.T) {
	mockAdminRepository := mockAdminRepo{}
	scenarios := []Scenario{
		{
			TestName: "when email is valid",
			Email: "admin@email.com",
			Token: "valid_token",
			Err: nil,
			findByEmailFunc: func(email string) (*entities.Admin, error) {
				return &entities.Admin{
					Id:        0,
					Email:     "admin@email.com",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}, nil
			},
		},
		{
			TestName: "when email is invalid",
			Email: "invalid@email.com",
			Token: "",
			Err: errors.New("admin not found"),
			findByEmailFunc: func(email string) (*entities.Admin, error) {
				return nil, errors.New("admin not found")
			},
		},
		{
			TestName: "when email isn't passed",
			Email: "",
			Token: "",
			Err: errors.New("email must be provided"),
			findByEmailFunc: func(email string) (*entities.Admin, error) {
				return nil, errors.New("admin not found")
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			mockAdminRepository.findByEmailFunc = scenario.findByEmailFunc
			repositories.AdminRepo = mockAdminRepository
			token, err := Authenticate(scenario.Email)
			assert.EqualValues(t, scenario.Token, token)
			assert.EqualValues(t, scenario.Err, err)
		})
	}
}