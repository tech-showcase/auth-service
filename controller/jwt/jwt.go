package jwt

import (
	"errors"
	"github.com/tech-showcase/auth-service/helper"
	"github.com/tech-showcase/auth-service/model"
	"net/http"
	"strings"
	"time"
)

type (
	RegisterRequest struct {
		model.UserData
	}

	RegisterResponse struct {
		model.UserCredential
	}

	LoginRequest struct {
		Phone string `json:"phone"`
		model.UserCredential
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	AuthHeader struct {
		Authorization string `header:"Authorization"`
	}
)

const (
	PrivateClaimsContextKey = "claims"
)

func Register(request RegisterRequest, userRepo model.UserRepo) (response RegisterResponse, statusCode int, err error) {
	password := helper.Generate4CharsPassword(request.Phone)
	userCredential := model.UserCredential{
		Password: password,
	}

	user := model.User{
		UserData:       request.UserData,
		UserCredential: userCredential,
	}
	_, err = userRepo.AddUser(user)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	response = RegisterResponse{
		UserCredential: userCredential,
	}
	statusCode = http.StatusOK
	return
}

func Login(request LoginRequest, userRepo model.UserRepo, authHelper helper.AuthHelper) (response LoginResponse, statusCode int, err error) {
	user := userRepo.GetUserByPhone(request.Phone)
	if user.Password != request.Password {
		statusCode = http.StatusUnauthorized
		err = errors.New("phone and password is not matched")
		return
	}

	privateClaims := helper.PrivateClaims{
		UserData:  user.UserData,
		Timestamp: time.Now().Unix(),
	}
	token, err := authHelper.GenerateToken(privateClaims, user.Password)
	if err != nil {
		statusCode = http.StatusInternalServerError
		err = errors.New("failed to generate token")
		return
	}

	response = LoginResponse{
		Token: token,
	}
	statusCode = http.StatusOK
	return
}

func AuthenticateJWT(header AuthHeader, authHelper helper.AuthHelper) (privateClaims helper.PrivateClaims, statusCode int, err error) {
	token := strings.TrimPrefix(header.Authorization, "Bearer ")
	token = strings.TrimSpace(token)

	privateClaims, err = authHelper.ParseToken(token)
	if err != nil {
		statusCode = http.StatusUnauthorized
		err = errors.New("token is invalid")
		return
	}

	user := model.UserRepoInstance.GetUserByPhone(privateClaims.Phone)
	isValid := authHelper.ValidateToken(token, user.Password)
	if !isValid {
		statusCode = http.StatusUnauthorized
		err = errors.New("token is invalid")
		return
	}

	privateClaims, err = authHelper.ParseToken(token)
	if err != nil {
		statusCode = http.StatusUnauthorized
		err = errors.New("token is invalid")
		return
	}

	statusCode = http.StatusOK
	return
}
