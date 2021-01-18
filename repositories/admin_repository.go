package repositories

import "github.com/diogoqds/routes-challenge-api/entities"

type FindAdmin interface {
	FindByEmail(email string) (*entities.Admin, error)
}

type AdminRepository struct {}

func (a AdminRepository) FindByEmail(email string) (*entities.Admin, error) {
	return nil, nil

}

var (
	AdminRepo FindAdmin = AdminRepository{}
)
