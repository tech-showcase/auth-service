package api

import (
	"github.com/gin-gonic/gin"
	"go-simple-app/controller"
)

func RegisterAuthAPI(router *gin.Engine) {
	router.POST("/api/login", controller.Login)
}
