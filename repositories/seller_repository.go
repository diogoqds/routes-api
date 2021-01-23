package repositories

import (
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	"log"
	"time"
)

type CreateSeller interface {
	Create(name string, email string) (*entities.Seller, error)
}

type SellerRepository struct {
	CreateSeller CreateSeller
}

type createSellerImplementation struct{}

func (c createSellerImplementation) Create(name string, email string) (*entities.Seller, error) {
	seller := entities.Seller{
		Name:      name,
		Email:     email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
	sqlInsert := "INSERT INTO sellers (name, email) VALUES ($1, $2) RETURNING id"
	var id int

	err := infra.DB.QueryRow(sqlInsert, seller.Name, seller.Email).Scan(&id)

	if err != nil {
		log.Println("Error saving the seller: " + err.Error())
		return nil, err
	}

	seller.Id = id
	return &seller, nil
}

var (
	SellerRepo = SellerRepository{
		CreateSeller: createSellerImplementation{},
	}
)
