package api

import (
	"github.com/gin-gonic/gin"
	supportPresenter "github.com/tech-showcase/auth-service/presenter/support"
)

func RegisterSupportAPI(router *gin.Engine) {
	router.GET("/health-check", supportPresenter.HealthCheck)
}
