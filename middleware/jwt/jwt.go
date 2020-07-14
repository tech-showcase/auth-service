package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller/jwt"
	"github.com/tech-showcase/auth-service/helper"
	"net/http"
	"strings"
)

func AuthenticationMiddleware(ctx *gin.Context) {
	authHeader := jwt.AuthHeader{}
	if err := ctx.ShouldBindHeader(&authHeader); err != nil || authHeader.Authorization == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "request should contains authorization header"})
		ctx.Abort()
		return
	}

	if !strings.Contains(authHeader.Authorization, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "token should contains Bearer"})
		ctx.Abort()
		return
	}

	privateClaims, statusCode, err := jwt.AuthenticateJWT(authHeader, helper.NewAuthHelper())
	if err != nil {
		ctx.JSON(statusCode, map[string]string{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx = setPrivateClaimsToContext(ctx, privateClaims)
	ctx.Next()
}

func setPrivateClaimsToContext(ctx *gin.Context, privateClaims helper.PrivateClaims) *gin.Context {
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys[jwt.PrivateClaimsContextKey] = privateClaims

	return ctx
}
