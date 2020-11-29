package api

import (
	"github.com/gin-gonic/gin"
	oauth2Middleware "github.com/tech-showcase/auth-service/middleware/oauth2"
	oauth2Presenter "github.com/tech-showcase/auth-service/presenter/oauth2"
)

func RegisterOAuth2API(router *gin.Engine) {
	oauth2Route := router.Group("/oauth2")
	oauth2Route.GET("/login", oauth2Presenter.GetLogin)
	oauth2Route.POST("/login", oauth2Presenter.PostLogin)
	oauth2Route.POST("/token", oauth2Presenter.Token)
	oauth2Route.Use(oauth2Middleware.SessionDataMiddleware, oauth2Middleware.MustBeLoggedInMiddleware)
	{
		oauth2Route.GET("/authorize", oauth2Presenter.Authorize)
		oauth2Route.GET("/authorization", oauth2Presenter.GetAuthorization)
		oauth2Route.POST("/authorize", oauth2Presenter.Authorize)
	}
}
