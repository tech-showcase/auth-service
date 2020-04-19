package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/helper"
	"github.com/tech-showcase/auth-service/model"
	"github.com/tech-showcase/auth-service/presenter"
	"github.com/tech-showcase/auth-service/singleton"
	"net/http"
	"strconv"
	"time"
)

const (
	ClaimsContextKey = "claims"
)

func Register(ctx *gin.Context) {
	registerRequest := presenter.RegisterRequestStruct{}
	if err := ctx.ShouldBind(&registerRequest); err == nil {
		password := helper.Generate4CharsPassword(registerRequest.Phone)
		userCredential := model.UserCredential{
			Password: password,
		}
		registerResponse := presenter.RegisterResponseStruct{
			UserCredential: userCredential,
		}

		userData := model.User{
			UserData:       registerRequest.UserData,
			UserCredential: userCredential,
		}
		singleton.UsersRepo.AddOrUpdateUser(userData)

		ctx.JSON(http.StatusOK, registerResponse)
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "request body should contains JSON"})
	}
}

func Login(ctx *gin.Context) {
	loginRequest := presenter.LoginRequestStruct{}
	if err := ctx.ShouldBind(&loginRequest); err == nil {
		userData := singleton.UsersRepo.GetUserByPhone(loginRequest.Phone)
		if userData.Password != loginRequest.Password {
			ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "user and password is not correct"})
			return
		}

		privateClaims := presenter.PrivateClaims{
			UserData:  userData.UserData,
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		}
		authHelper := helper.NewAuthBlueprint()
		token, err := authHelper.GenerateToken(privateClaims, userData.Password)
		loginResponse := presenter.LoginResponseStruct{
			Token: token,
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to generate token"})
		} else {
			ctx.JSON(http.StatusOK, loginResponse)
		}
	}
}

func GetActiveUser(ctx *gin.Context) {
	claims := ctx.Keys[ClaimsContextKey]
	ctx.JSON(http.StatusOK, claims)
}
