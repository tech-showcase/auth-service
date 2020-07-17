package oauth2

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller/oauth2"
	"net/http"
)

func Authorize(ctx *gin.Context) {
	err := oauth2.Authorize(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
}

func Token(ctx *gin.Context) {
	err := oauth2.Token(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
}
