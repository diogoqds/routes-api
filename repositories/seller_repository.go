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

type ListSellers interface {
	FindAll() ([]entities.Seller, error)
}

type SellerRepository struct {
	CreateSeller CreateSeller
	ListSellers  ListSellers
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

func (c createSellerImplementation) FindAll() ([]entities.Seller, error) {
	var err error
	sellers := make([]entities.Seller, 0)
	query := "SELECT * FROM sellers WHERE deleted_at IS NULL"
	err = infra.DB.Select(&sellers, query)

	if err != nil {
		return nil, err
	}

	return sellers, nil
}

var (
	SellerRepo = SellerRepository{
		CreateSeller: createSellerImplementation{},
		ListSellers:  createSellerImplementation{},
	}
)
