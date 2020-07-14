package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller"
	"github.com/tech-showcase/auth-service/helper"
	"net/http"
	"strings"
)

func JWTAuthenticationMiddleware(ctx *gin.Context) {
	authHeader := controller.AuthHeader{}
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

	privateClaims, statusCode, err := controller.AuthenticateJWT(authHeader, helper.NewAuthHelper())
	if err != nil {
		ctx.JSON(statusCode, map[string]string{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx = setJWTPrivateClaimsToContext(ctx, privateClaims)
	ctx.Next()
}

func setJWTPrivateClaimsToContext(ctx *gin.Context, privateClaims helper.PrivateClaims) *gin.Context {
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys[controller.JWTPrivateClaimsContextKey] = privateClaims

	return ctx
}
