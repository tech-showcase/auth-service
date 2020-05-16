package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller"
	"github.com/tech-showcase/auth-service/middleware"
	"io"
	"os"
	"strconv"
)

func ActivateHTTP(port int) {
	setupHTTPLogger()

	router := gin.Default()

	RegisterAuthAPI(router)

	address := ":" + strconv.Itoa(port)
	router.Run(address)
}

func setupHTTPLogger() {
	f, _ := os.Create("http.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func RegisterAuthAPI(router *gin.Engine) {
	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)
	router.Use(middleware.JWTAuthenticationMiddleware)
	{
		router.GET("/api/user", controller.GetActiveUser)
	}
}
