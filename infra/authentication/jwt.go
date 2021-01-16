package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type JsonWebToken interface {
	Encode(body map[string]interface{}) (string, error)
}

type jsonWebToken struct{}

func (j jsonWebToken) Encode(body map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(body))
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return signedToken, err
}

var (
	Jwt JsonWebToken = jsonWebToken{}
)
