package usecases

import (
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"log"
)

type ListSellerClients interface {
	FindAll(sellerId int) ([]entities.Client, error)
}

type ListSellerClientsUseCase struct {
}

func (l ListSellerClientsUseCase) FindAll(sellerId int) ([]entities.Client, error) {
	route, err := repositories.SellerRepo.FinderSellerRoute.FindRoute(sellerId)

	if err != nil {
		log.Println("error fetching seller route", err.Error())
		return nil, err
	}

	clients, err := repositories.ClientRepo.ClientsFinderByRouteId.FindAllByRouteId(route.Id)

	if err != nil {
		log.Println("error fetching clients by route", err.Error())
		return nil, err
	}

	return clients, nil
}

var (
	ListSellerClientsService ListSellerClients = ListSellerClientsUseCase{}
)
