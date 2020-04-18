package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller"
	"github.com/tech-showcase/auth-service/middleware"
	"strconv"
)

func ActivateHTTP(port int) {
	router := gin.Default()

	RegisterAuthAPI(router)

	portStr := ":" + strconv.Itoa(port)
	router.Run(portStr)
}


func RegisterAuthAPI(router *gin.Engine) {
	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)
	router.Use(middleware.JWTAuthenticationMiddleware)
	{
		router.GET("/api/user", controller.GetActiveUser)
	}
}
