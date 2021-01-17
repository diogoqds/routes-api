package infra

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	RealEncodeImplementation func(body map[string]interface{}) (string, error)
)

func init() {
	RealEncodeImplementation = Jwt.Encoder.Encode
}

type jsonWebTokenEncoderMock struct {
	encodeFunc func(body map[string]interface{}) (string, error)
}

func (mock jsonWebTokenEncoderMock) Encode(body map[string]interface{}) (string, error) {
	return mock.encodeFunc(body)
}

func cleanMocks() {
	jwtEncoderMock := jsonWebTokenEncoderMock{}
	jwtEncoderMock.encodeFunc = RealEncodeImplementation

	Jwt.Encoder = jwtEncoderMock
}

func TestJsonWebToken_EncodeSuccess(t *testing.T) {
	cleanMocks()
	body := map[string]interface{}{
		"id": 1,
	}

	token, err := Jwt.Encoder.Encode(body)

	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestJsonWebToken_EncodeError(t *testing.T) {
	cleanMocks()
	jwtEncoderMock := jsonWebTokenEncoderMock{}
	jwtEncoderMock.encodeFunc = func(body map[string]interface{}) (string, error) {
		return "", errors.New("error generating the JWT token")
	}

	Jwt.Encoder = jwtEncoderMock

	body := map[string]interface{}{
		"id": 1,
	}
	token, err := Jwt.Encoder.Encode(body)

	assert.Errorf(t, err, "error generating the JWT token")
	assert.Empty(t, token)
}

func TestJsonWebToken_DecodeSuccess(t *testing.T) {
	cleanMocks()

	body := map[string]interface{}{
		"id": 1,
	}

	token, _ := Jwt.Encoder.Encode(body)

	result, err := Jwt.Decoder.Decode(token)

	assert.Nil(t, err)
	assert.EqualValues(t, 1, result["id"])
}

func TestJsonWebToken_DecodeError(t *testing.T) {
	cleanMocks()
	result, err := Jwt.Decoder.Decode("invalid token")

	assert.EqualValues(t, "token contains an invalid number of segments", err.Error())
	assert.EqualValues(t, jwt.MapClaims(nil), result)
}
