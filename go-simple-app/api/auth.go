package api

import (
	"github.com/gin-gonic/gin"
	"go-simple-app/controller"
	"go-simple-app/middleware"
)

func RegisterAuthAPI(router *gin.Engine) {
	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)
	router.Use(middleware.JWTAuthenticationMiddleware)
	{
		router.POST("/api/user", controller.GetActiveUser)
	}
}
