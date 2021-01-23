package repositories

import (
	"database/sql"
	"errors"
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-sqlxmock"
	"regexp"
	"testing"
)

func TestFindByEmail_Success(t *testing.T) {
	var db *sqlx.DB
	var err error
	var mock sqlmock.Sqlmock

	db, mock, err = sqlmock.Newx()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}

	infra.DB = db
	rows := mock.NewRows([]string{"id", "email"}).AddRow(1, "admin@email.com")

	sql := "SELECT id, email FROM admins WHERE email = $1"
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WillReturnRows(rows)
	admin, err := AdminRepo.FinderByEmail.FindByEmail("admin@email.com")

	assert.EqualValues(t, 1, admin.Id)
	assert.EqualValues(t, "admin@email.com", admin.Email)
	assert.Nil(t, err)
}

func TestFindByEmail_Error(t *testing.T) {
	scenarios := []struct {
		TestName     string
		Error        error
		ErrorMessage string
	}{
		{
			TestName:     "when there is no admin",
			Error:        sql.ErrNoRows,
			ErrorMessage: "sql: no rows in result set",
		},
		{
			TestName:     "when a generic error happens",
			Error:        errors.New("generic error"),
			ErrorMessage: "generic error",
		},
	}
	var db *sqlx.DB
	var err error
	var mock sqlmock.Sqlmock

	db, mock, err = sqlmock.Newx()
	infra.DB = db
	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

			}

			query := "SELECT id, email FROM admins WHERE email = $1"
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(scenario.Error)
			admin, err := AdminRepo.FinderByEmail.FindByEmail("admin@email.com")

			assert.EqualValues(t, scenario.ErrorMessage, err.Error())
			assert.Nil(t, admin)
		})
	}
}
