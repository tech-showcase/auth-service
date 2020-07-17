package oauth2

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller/oauth2"
)

func Authorize(ctx *gin.Context) {
	oauth2.Authorize(ctx.Writer, ctx.Request)
}

func Token(ctx *gin.Context) {
	oauth2.Token(ctx.Writer, ctx.Request)
}
