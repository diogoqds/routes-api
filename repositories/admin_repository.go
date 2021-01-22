package repositories

import (
	"database/sql"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	"log"
)

type FindAdmin interface {
	FindByEmail(email string) (*entities.Admin, error)
	FindById(id int64) (*entities.Admin, error)
}

type AdminRepository struct{}

func (a AdminRepository) FindByEmail(email string) (*entities.Admin, error) {
	var admin entities.Admin

	err := infra.DB.Get(&admin, "SELECT id, email FROM admins WHERE email = $1", email)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no admin with email %d\n", email)
		return nil, err
	case err != nil:
		log.Fatalf("query error: %v\n", err)
		return nil, err
	default:
		log.Printf("admin found it")
		return &admin, nil
	}
}

func (a AdminRepository) FindById(id int64) (*entities.Admin, error) {
	var admin entities.Admin

	err := infra.DB.Get(&admin, "SELECT id, email FROM admins WHERE id = $1", id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no admin with id %d\n", id)
		return nil, err
	case err != nil:
		log.Fatalf("query error: %v\n", err)
		return nil, err
	default:
		log.Printf("admin found it")
		return &admin, nil
	}
}

var (
	AdminRepo FindAdmin = AdminRepository{}
)
