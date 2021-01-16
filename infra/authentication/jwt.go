package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
}

type JsonWebTokenEncoder interface {
	Encode(body map[string]interface{}) (string, error)
}

type JsonWebTokenDecoder interface {
	Decode(token string) (jwt.MapClaims, error)
}

type JsonWebToken struct {
	Encoder JsonWebTokenEncoder
	Decoder JsonWebTokenDecoder
}

type jsonWebTokenImplementation struct {
}

func (j jsonWebTokenImplementation) Encode(body map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(body))
	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return signedToken, err
}

func (j jsonWebTokenImplementation) Decode(token string) (jwt.MapClaims, error) {

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

var (
	Jwt = JsonWebToken{
		Encoder: jsonWebTokenImplementation{},
		Decoder: jsonWebTokenImplementation{},
	}
)
