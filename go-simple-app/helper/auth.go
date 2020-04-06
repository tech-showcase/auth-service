package helper

import (
	"github.com/dgrijalva/jwt-go"
	"go-simple-app/presenter"
)

type (
	AuthClaims struct {
		presenter.PrivateClaims
		jwt.StandardClaims
	}
	AuthBlueprint struct{}
	AuthInterface interface {
		GenerateToken(presenter.PrivateClaims, string) (string, error)
	}
)

func NewAuthBlueprint() AuthInterface {
	var instance AuthBlueprint

	return &instance
}

func (instance *AuthBlueprint) GenerateToken(privateClaims presenter.PrivateClaims, key string) (string, error) {
	claims := AuthClaims{
		PrivateClaims:  privateClaims,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
