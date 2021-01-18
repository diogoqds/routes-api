package usecases

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

func Authenticate(email string) (string, error) {

	if email == "" {
		return "", errors.New("email must be provided")
	}

	_, err := repositories.AdminRepo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	return "valid_token", nil
}