package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller"
	"github.com/tech-showcase/auth-service/global"
	"github.com/tech-showcase/auth-service/middleware"
	"go.elastic.co/apm/module/apmgin"
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
	authRoute := router.Group("/api")
	authRoute.POST("/register", controller.Register)
	authRoute.POST("/login", controller.Login)
	authRoute.Use(middleware.JWTAuthenticationMiddleware)
	{
		authRoute.GET("/user", controller.GetActiveUser)
	}
}
