package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller/jwt"
	"github.com/tech-showcase/auth-service/helper"
	"github.com/tech-showcase/auth-service/model"
	"net/http"
)

func Register(ctx *gin.Context) {
	registerRequest := jwt.RegisterRequest{}
	if err := ctx.ShouldBind(&registerRequest); err == nil {
		registerResponse, statusCode, err := jwt.Register(registerRequest, model.UserRepoInstance)
		if err != nil {
			ctx.JSON(statusCode, map[string]string{"message": err.Error()})
		} else {
			ctx.JSON(statusCode, registerResponse)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "request body format is wrong"})
	}
}

func Login(ctx *gin.Context) {
	loginRequest := jwt.LoginRequest{}
	if err := ctx.ShouldBind(&loginRequest); err == nil {
		loginResponse, statusCode, err := jwt.Login(loginRequest, model.UserRepoInstance, helper.NewAuthHelper())
		if err != nil {
			ctx.JSON(statusCode, map[string]string{"message": err.Error()})
		} else {
			ctx.JSON(statusCode, loginResponse)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "request body format is wrong"})
	}
}

func GetActiveUser(ctx *gin.Context) {
	claims := ctx.Keys[jwt.PrivateClaimsContextKey]
	ctx.JSON(http.StatusOK, claims)
}
