package oauth2

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller/oauth2"
	"github.com/tech-showcase/auth-service/helper"
	"net/http"
)

func SessionDataMiddleware(ctx *gin.Context) {
	sessionData, err := helper.NewSessionData(helper.SessionStorageInstance, helper.OAuth2SessionName, ctx.Request, ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	oauth2.SaveAuthorizationData(ctx.Request, sessionData)
	ctx.Next()
}

func MustBeLoggedInMiddleware(ctx *gin.Context) {
	sessionData, err := helper.NewSessionData(helper.SessionStorageInstance, helper.OAuth2SessionName, ctx.Request, ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	if !oauth2.IsLoggedIn(sessionData) {
		ctx.Redirect(http.StatusFound, helper.LoginUrl)
		ctx.Abort()
	}

	ctx.Next()
}
