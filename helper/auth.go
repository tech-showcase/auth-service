package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/tech-showcase/auth-service/model"
	"github.com/tech-showcase/auth-service/presenter"
	"strings"
)

type (
	authClaims struct {
		presenter.PrivateClaims
		jwt.StandardClaims
	}
	authBlueprint struct{}
	AuthInterface interface {
		GenerateToken(presenter.PrivateClaims, string) (string, error)
		ParseAndValidateToken(string, string) (presenter.PrivateClaims, error)
		ParseToken(tokenStr string) (presenter.PrivateClaims, error)
	}
)

func NewAuthBlueprint() AuthInterface {
	var instance authBlueprint

	return &instance
}

func (instance *authBlueprint) GenerateToken(privateClaims presenter.PrivateClaims, key string) (string, error) {
	claims := authClaims{
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

func (instance *authBlueprint) ParseAndValidateToken(tokenStr, key string) (presenter.PrivateClaims, error) {
	tokenStr = strings.TrimSpace(tokenStr)

	token, err := jwt.ParseWithClaims(tokenStr, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return presenter.PrivateClaims{}, err
	}

	if claims, ok := token.Claims.(*authClaims); ok && token.Valid {
		return claims.PrivateClaims, nil
	} else {
		return presenter.PrivateClaims{}, errors.New("token is invalid")
	}
}

func (instance *authBlueprint) ParseToken(tokenStr string) (presenter.PrivateClaims, error) {
	tokenStr = strings.TrimSpace(tokenStr)

	token, _ := jwt.Parse(tokenStr, nil)

	if claimsMap, ok := token.Claims.(jwt.MapClaims); ok {
		userData := model.UserData{
			Username:  claimsMap["name"].(string),
			Phone: claimsMap["phone"].(string),
			Email: claimsMap["email"].(string),
		}
		timestamp := int64(claimsMap["timestamp"].(float64))
		privateClaims := presenter.PrivateClaims{
			UserData:  userData,
			Timestamp: timestamp,
		}
		return privateClaims, nil
	} else {
		return presenter.PrivateClaims{}, errors.New("token is invalid")
	}
}
