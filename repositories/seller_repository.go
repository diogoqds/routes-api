package repositories

import (
	"log"
	"time"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
)

type CreateSeller interface {
	Create(name string, email string) (*entities.Seller, error)
}

type ListSellers interface {
	FindAll() ([]entities.Seller, error)
}

type DeleteSeller interface {
	Delete(id int) (bool, error)
}

type FinderSellerRoute interface {
	FindRoute(sellerId int) (*entities.Route, error)
}

type SellerRepository struct {
	CreateSeller      CreateSeller
	ListSellers       ListSellers
	DeleteSeller      DeleteSeller
	FinderSellerRoute FinderSellerRoute
}

type sellerRepository struct{}

func (c sellerRepository) Create(name string, email string) (*entities.Seller, error) {
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

func (c sellerRepository) FindAll() ([]entities.Seller, error) {
	var err error
	sellers := make([]entities.Seller, 0)
	query := "SELECT * FROM sellers WHERE deleted_at IS NULL"
	err = infra.DB.Select(&sellers, query)

	if err != nil {
		return nil, err
	}

	return sellers, nil
}

func (c sellerRepository) Delete(id int) (bool, error) {
	var sellerId int

	query := "UPDATE sellers SET deleted_at = NOW() WHERE id = $1 RETURNING id"

	err := infra.DB.QueryRow(query, id).Scan(&sellerId)

	if err != nil {
		log.Println("Error deleting the seller: " + err.Error())
		return false, err
	}

	return sellerId > 0, nil
}

func (c sellerRepository) FindRoute(sellerId int) (*entities.Route, error) {
	var route entities.Route
	query := "SELECT routes.id as id FROM routes JOIN sellers ON routes.seller_id = sellers.id WHERE seller_id = $1"

	err := infra.DB.Get(&route, query, sellerId)

	if err != nil {
		return nil, err
	}

	return &route, nil
}

var (
	SellerRepo = SellerRepository{
		CreateSeller:      sellerRepository{},
		ListSellers:       sellerRepository{},
		DeleteSeller:      sellerRepository{},
		FinderSellerRoute: sellerRepository{},
	}
)
