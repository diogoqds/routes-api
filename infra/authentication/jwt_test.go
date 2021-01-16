package authentication

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type jsonWebTokenMock struct {
	encodeFunc func(body map[string]interface{}) (string, error)
}

func (mock jsonWebTokenMock) Encode(body map[string]interface{}) (string, error) {
	return mock.encodeFunc(body)
}

func TestJsonWebToken_EncodeSuccess(t *testing.T) {
	body := map[string]interface{}{
		"id": 1,
	}

	token, err := Jwt.Encode(body)

	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestJsonWebToken_EncodeError(t *testing.T) {
	jwtMock := jsonWebTokenMock{}
	jwtMock.encodeFunc = func(body map[string]interface{}) (string, error) {
		return "", errors.New("error generating the JWT token")
	}

	Jwt = jwtMock

	body := map[string]interface{}{
		"id": 1,
	}
	token, err := Jwt.Encode(body)

	assert.Errorf(t, err, "error generating the JWT token")
	assert.Empty(t, token)
}
