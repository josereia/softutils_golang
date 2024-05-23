package auth

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type AuthHeaderUtil struct {
	Context *gin.Context
}

type BasicCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (util *AuthHeaderUtil) Basic() (BasicCredentials, error) {
	header := util.Context.Request.Header.Get("Authorization")
	if header == "" {
		return BasicCredentials{}, errors.New("authorization header is required")
	}

	encoded := strings.Split(header, " ")
	if len(encoded) != 2 {
		return BasicCredentials{}, errors.New("authorization header is malformed")
	}

	decoded, err := base64.StdEncoding.DecodeString(encoded[1])
	if err != nil {
		return BasicCredentials{}, errors.Wrap(err, "authorization header is malformed")

	}

	credentials := strings.Split(string(decoded), ":")
	if len(credentials) != 2 {
		return BasicCredentials{}, errors.New("authorization header is malformed")
	}

	response := BasicCredentials{
		Username: credentials[0],
		Password: credentials[1],
	}

	return response, nil
}

func (util *AuthHeaderUtil) Bearer() (string, error) {
	header := util.Context.Request.Header.Get("Authorization")

	token := strings.Split(header, " ")
	if len(token) != 2 {
		return "", errors.New("authorization header is malformed")
	}

	return token[1], nil
}
