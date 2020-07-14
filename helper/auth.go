package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/tech-showcase/auth-service/model"
	"strings"
)

type (
	PrivateClaims struct {
		model.UserData
		Timestamp int64 `json:"timestamp"`
	}
	authClaims struct {
		PrivateClaims
		jwt.StandardClaims
	}
	authHelper struct{}
	AuthHelper interface {
		GenerateToken(privateClaims PrivateClaims, key string) (token string, err error)
		ValidateToken(token, key string) (isValid bool)
		ParseToken(token string) (privateClaims PrivateClaims, err error)
	}
)

func NewAuthHelper() AuthHelper {
	var instance authHelper

	return &instance
}

func (instance *authHelper) GenerateToken(privateClaims PrivateClaims, key string) (token string, err error) {
	claims := authClaims{
		PrivateClaims:  privateClaims,
		StandardClaims: jwt.StandardClaims{},
	}

	jwToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwToken.SignedString([]byte(key))
	return
}

func (instance *authHelper) ValidateToken(token, key string) (isValid bool) {
	token = strings.TrimSpace(token)
	jwToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return
	}

	isValid = jwToken.Valid
	return
}

func (instance *authHelper) ParseToken(token string) (privateClaims PrivateClaims, err error) {
	token = strings.TrimSpace(token)
	jwToken, _ := jwt.Parse(token, nil)

	if claimsMap, ok := jwToken.Claims.(jwt.MapClaims); ok {
		userData := model.UserData{
			Username: claimsMap["username"].(string),
			Phone:    claimsMap["phone"].(string),
			Email:    claimsMap["email"].(string),
			IsActive: claimsMap["is_active"].(bool),
		}
		timestamp := int64(claimsMap["timestamp"].(float64))
		privateClaims = PrivateClaims{
			UserData:  userData,
			Timestamp: timestamp,
		}
		return
	} else {
		err = errors.New("token is invalid")
		return
	}
}
