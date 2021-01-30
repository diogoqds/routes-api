package usecases

import (
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type DeleteRoute interface {
	Delete(id int) (bool, error)
}

type DeleteRouteUseCase struct {
}

func (d DeleteRouteUseCase) Delete(id int) (bool, error) {
	deleted, err := repositories.RouteRepo.RouteEraser.Delete(id)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

var (
	DeleteRouteService DeleteRoute = DeleteRouteUseCase{}
)
