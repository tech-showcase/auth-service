package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/middleware"
	"github.com/tech-showcase/auth-service/presenter"
)

func RegisterJWTAPI(router *gin.Engine) {
	authRoute := router.Group("/api")
	authRoute.POST("/register", presenter.Register)
	authRoute.POST("/login", presenter.Login)
	authRoute.Use(middleware.JWTAuthenticationMiddleware)
	{
		authRoute.GET("/user", presenter.GetActiveUser)
	}
}
