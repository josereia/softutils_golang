package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type RefreshClaims struct{ jwt.StandardClaims }
type AccessClaims struct {
	jwt.StandardClaims
	Id    string `json:"id"`
	Roles []uint `json:"roles"`
}

type AuthTokenUtil struct {
	Secret string `json:"secret"`
}

// Generates an access token based on claims
func (util *AuthTokenUtil) GenAccessToken(claims AccessClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(util.Secret))
	if err != nil {
		return "", errors.Wrap(err, "error on sign access token")
	}

	return signed, nil
}

// Generates a refresh token based on claims
func (util *AuthTokenUtil) GenRefreshToken(
	claims RefreshClaims,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(util.Secret))
	if err != nil {
		return "", errors.Wrap(err, "error on sign refresh token")
	}

	return signed, nil
}

// Decodes the access token and returns the claims
func (util *AuthTokenUtil) GetAccessToken(token string) (AccessClaims, error) {
	parsed, err := jwt.ParseWithClaims(
		token,
		&AccessClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(util.Secret), nil
		},
	)
	if err != nil {
		return AccessClaims{}, errors.Wrap(err, "error on parse access token")
	}

	claims := parsed.Claims.(*AccessClaims)
	return *claims, nil
}

// Decodes the refresh token and returns the claims
func (util *AuthTokenUtil) GetRefreshToken(
	token string,
) (RefreshClaims, error) {
	parsed, err := jwt.ParseWithClaims(
		token,
		&RefreshClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(util.Secret), nil
		},
	)
	if err != nil {
		return RefreshClaims{}, errors.Wrap(err, "error on parse refresh token")
	}

	claims := parsed.Claims.(*RefreshClaims)
	return *claims, nil
}

// Checks the validity of the access token
func (claims *AccessClaims) Verify() (bool, error) {
	stdErr := claims.StandardClaims.Valid()
	err := claims.Valid()

	if stdErr != nil && err == nil {
		return false, errors.Wrap(stdErr, "invalid access token")
	}

	return true, nil
}

// Checks the validity of the refresh token
func (claims *RefreshClaims) Verify() (bool, error) {
	err := claims.Valid()
	if err != nil {
		return false, errors.Wrap(err, "invalid refresh token")
	}

	return true, nil
}
