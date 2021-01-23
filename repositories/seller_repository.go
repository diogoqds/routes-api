package repositories

import (
	"fmt"
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
	sqlInsert := "INSERT INTO sellers (name, email) VALUES ($1, $2)"
	result, err := infra.DB.Exec(sqlInsert, seller.Name, seller.Email)

	if err != nil {
		log.Println("Error saving the seller: " + err.Error())
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting the last insert id for new seller: " + err.Error())
		return nil, err
	}
	seller.Id = int(id)
	return &seller, nil
}

var (
	SellerRepo = SellerRepository{
		CreateSeller: createSellerImplementation{},
	}
)
