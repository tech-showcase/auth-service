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
	setupAPMAgent(router)

	RegisterAuthAPI(router)

	address := ":" + strconv.Itoa(port)
	router.Run(address)
}

func setupHTTPLogger() {
	f, err := os.OpenFile(global.Configuration.Log.FilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func setupAPMAgent(router *gin.Engine) {
	router.Use(apmgin.Middleware(router))
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
