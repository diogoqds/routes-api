package repositories

import (
	"database/sql"

	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	"log"
)

type FinderAdminByEmail interface {
	FindByEmail(email string) (*entities.Admin, error)
}

type FinderAdminById interface {
	FindById(id int64) (*entities.Admin, error)
}

type AdminRepository struct {
	FinderByEmail FinderAdminByEmail
	FinderById    FinderAdminById
}

type adminRepositoryImplementation struct{}

func (a adminRepositoryImplementation) FindByEmail(email string) (*entities.Admin, error) {
	var admin entities.Admin

	err := infra.DB.Get(&admin, "SELECT id, email FROM admins WHERE email = $1", email)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no admin with email %s\n", email)
		return nil, err
	case err != nil:
		log.Printf("query error: %v\n", err)
		return nil, err
	default:
		log.Printf("admin found it")
		return &admin, nil
	}
}

func (a adminRepositoryImplementation) FindById(id int64) (*entities.Admin, error) {
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
	AdminRepo = AdminRepository{
		FinderByEmail: adminRepositoryImplementation{},
		FinderById:    adminRepositoryImplementation{},
	}
)
