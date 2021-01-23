package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/diogoqds/routes-challenge-api/entities"
	"github.com/diogoqds/routes-challenge-api/infra"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"net/http"
	"time"
)

type mockJwtDecoder struct {
	decodeFunc func(token string) (jwt.MapClaims, error)
}

type mockAdminRepo struct {
	findByIdFunc func(id int64) (*entities.Admin, error)
}

func (mock mockJwtDecoder) Decode(token string) (jwt.MapClaims, error) {
	return mock.decodeFunc(token)
}

func (mock mockAdminRepo) FindById(id int64) (*entities.Admin, error) {
	return mock.findByIdFunc(id)
}

func setupAuth(request *http.Request) {
	request.Header.Set("Authorization", "Bearer token")

	mockJwt := mockJwtDecoder{decodeFunc: func(token string) (jwt.MapClaims, error) {
		return jwt.MapClaims{"id": float64(1)}, nil
	}}
	mockAdminRepository := mockAdminRepo{findByIdFunc: func(id int64) (*entities.Admin, error) {
		return &entities.Admin{
			Id:        1,
			Email:     "admin@email.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}, nil
	}}

	infra.Jwt.Decoder = mockJwt
	repositories.AdminRepo.FinderById = mockAdminRepository
}
