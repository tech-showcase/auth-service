package api

import (
	"github.com/gin-gonic/gin"
	oauth2Presenter "github.com/tech-showcase/auth-service/presenter/oauth2"
)

func RegisterOAuth2API(router *gin.Engine) {
	authRoute := router.Group("/oauth2")
	authRoute.GET("/authorize", oauth2Presenter.Authorize)
	authRoute.GET("/token", oauth2Presenter.Token)
}
