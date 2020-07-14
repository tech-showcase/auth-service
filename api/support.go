package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/presenter"
)

func RegisterSupportAPI(router *gin.Engine) {
	router.GET("/health-check", presenter.HealthCheck)
}
