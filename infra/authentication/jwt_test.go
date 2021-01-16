package authentication

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	RealEncodeImplementation func(body map[string]interface{}) (string, error)
)

func TestMain(t *testing.M) {
	RealEncodeImplementation = Jwt.Encoder.Encode
	t.Run()
}

type jsonWebTokenEncoderMock struct {
	encodeFunc func(body map[string]interface{}) (string, error)
}

func (mock jsonWebTokenEncoderMock) Encode(body map[string]interface{}) (string, error) {
	return mock.encodeFunc(body)
}

func TestJsonWebToken_EncodeSuccess(t *testing.T) {
	body := map[string]interface{}{
		"id": 1,
	}

	token, err := Jwt.Encoder.Encode(body)

	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestJsonWebToken_EncodeError(t *testing.T) {
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
	jwtEncoderMock := jsonWebTokenEncoderMock{}
	jwtEncoderMock.encodeFunc = RealEncodeImplementation

	body := map[string]interface{}{
		"id": 1,
	}

	Jwt.Encoder = jwtEncoderMock

	token, _ := Jwt.Encoder.Encode(body)

	result, err := Jwt.Decoder.Decode(token)

	assert.Nil(t, err)
	assert.EqualValues(t, 1, result["id"])
}

func TestJsonWebToken_DecodeError(t *testing.T) {
	result, err := Jwt.Decoder.Decode("invalid token")

	assert.EqualValues(t, "token contains an invalid number of segments", err.Error())
	assert.EqualValues(t, jwt.MapClaims(nil), result)
}
