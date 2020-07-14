package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller"
	"github.com/tech-showcase/auth-service/global"
	"github.com/tech-showcase/auth-service/helper"
	"net/http"
)

func Register(ctx *gin.Context) {
	registerRequest := controller.RegisterRequest{}
	if err := ctx.ShouldBind(&registerRequest); err == nil {
		registerResponse, statusCode, err := controller.Register(registerRequest, global.UsersRepo)
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
	loginRequest := controller.LoginRequest{}
	if err := ctx.ShouldBind(&loginRequest); err == nil {
		loginResponse, statusCode, err := controller.Login(loginRequest, global.UsersRepo, helper.NewAuthHelper())
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
	claims := ctx.Keys[controller.JWTPrivateClaimsContextKey]
	ctx.JSON(http.StatusOK, claims)
}
