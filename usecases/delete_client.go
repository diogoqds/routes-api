package usecases

import (
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type DeleteClient interface {
	Delete(id int) (bool, error)
}

type DeleteClientUseCase struct {
}

func (d DeleteClientUseCase) Delete(id int) (bool, error) {
	return repositories.ClientRepo.ClientEraser.Delete(id)
}

var (
	DeleteClientService DeleteClient = DeleteClientUseCase{}
)
