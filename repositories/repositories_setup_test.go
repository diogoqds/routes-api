package repositories

import (
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/jmoiron/sqlx"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
)

var (
	db   *sqlx.DB
	err  error
	mock sqlmock.Sqlmock
)

func setupTestDb() {
	db, mock, err = sqlmock.Newx()

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}

	infra.DB = db
}
