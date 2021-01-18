package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

func Authenticate(email string) (string, error) {

	if email == "" {
		return "", errors.New("email must be provided")
	}

	admin, err := repositories.AdminRepo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	token, err := infra.Jwt.Encoder.Encode(map[string]interface{}{ "id": admin.Id })

	if err != nil {
		return "", err
	}

	return token, nil
}