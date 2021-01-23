package repositories

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestFindByEmail_Success(t *testing.T) {
	setupDb()
	rows := mock.NewRows([]string{"id", "email"}).AddRow(1, "admin@email.com")

	query := "SELECT id, email FROM admins WHERE email = $1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
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
	setupDb()
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

func TestFindById_Success(t *testing.T) {
	setupDb()
	rows := mock.NewRows([]string{"id", "email"}).AddRow(1, "admin@email.com")

	sql := "SELECT id, email FROM admins WHERE id = $1"
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WillReturnRows(rows)
	admin, err := AdminRepo.FinderById.FindById(1)

	assert.EqualValues(t, 1, admin.Id)
	assert.EqualValues(t, "admin@email.com", admin.Email)
	assert.Nil(t, err)
}

func TestFindById_Error(t *testing.T) {
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
	setupDb()
	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

			}

			query := "SELECT id, email FROM admins WHERE id = $1"
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(scenario.Error)
			admin, err := AdminRepo.FinderById.FindById(1)

			assert.EqualValues(t, scenario.ErrorMessage, err.Error())
			assert.Nil(t, admin)
		})
	}
}
