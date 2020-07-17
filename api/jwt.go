package api

import (
	"github.com/gin-gonic/gin"
	jwtMiddleware "github.com/tech-showcase/auth-service/middleware/jwt"
	jwtPresenter "github.com/tech-showcase/auth-service/presenter/jwt"
)

func RegisterJWTAPI(router *gin.Engine) {
	authRoute := router.Group("/jwt")
	authRoute.POST("/register", jwtPresenter.Register)
	authRoute.POST("/login", jwtPresenter.Login)
	authRoute.Use(jwtMiddleware.AuthenticationMiddleware)
	{
		authRoute.GET("/user", jwtPresenter.GetActiveUser)
	}
}
