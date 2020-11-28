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
		userRepo := model.UserRepoInstance
		registerResponse, statusCode, err := jwt.Register(registerRequest, userRepo)
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
		userRepo := model.UserRepoInstance
		loginResponse, statusCode, err := jwt.Login(loginRequest, userRepo, helper.NewAuthHelper())
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
