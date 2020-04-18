package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/tech-showcase/auth-service/presenter"
	"strings"
)

type (
	AuthClaims struct {
		presenter.PrivateClaims
		jwt.StandardClaims
	}
	AuthBlueprint struct{}
	AuthInterface interface {
		GenerateToken(presenter.PrivateClaims, string) (string, error)
		ParseToken(string, string) (presenter.PrivateClaims, error)
		ParseTokenWithoutKey(tokenStr string) (presenter.PrivateClaims, error)
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
	tokenStr, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (instance *AuthBlueprint) ParseToken(tokenStr, key string) (presenter.PrivateClaims, error) {
	tokenStr = strings.TrimSpace(tokenStr)

	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return presenter.PrivateClaims{}, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.PrivateClaims, nil
	} else {
		return presenter.PrivateClaims{}, errors.New("token is invalid")
	}
}

func (instance *AuthBlueprint) ParseTokenWithoutKey(tokenStr string) (presenter.PrivateClaims, error) {
	tokenStr = strings.TrimSpace(tokenStr)

	token, _ := jwt.Parse(tokenStr, nil)

	if claimsMap, ok := token.Claims.(jwt.MapClaims); ok {
		privateClaims := presenter.PrivateClaims{
			Name:      claimsMap["name"].(string),
			Phone:     claimsMap["phone"].(string),
			Role:      claimsMap["role"].(string),
			Timestamp: claimsMap["timestamp"].(string),
		}
		return privateClaims, nil
	} else {
		return presenter.PrivateClaims{}, errors.New("token is invalid")
	}
}
