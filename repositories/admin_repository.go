package repositories

import (
	"context"
	"database/sql"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	"log"
)

type FindAdmin interface {
	FindByEmail(email string) (*entities.Admin, error)
}

type AdminRepository struct{}

func (a AdminRepository) FindByEmail(email string) (*entities.Admin, error) {
	var admin entities.Admin
	var ctx context.Context

	err := infra.DB.QueryRowContext(ctx, "SELECT id, email, created_at, updated_at FROM admins WHERE email=?", email).Scan(&admin)
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

var (
	AdminRepo FindAdmin = AdminRepository{}
)
