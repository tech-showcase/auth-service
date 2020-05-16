package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller"
	"github.com/tech-showcase/auth-service/helper"
	"github.com/tech-showcase/auth-service/presenter"
	"github.com/tech-showcase/auth-service/global"
	"net/http"
	"strings"
)

func JWTAuthenticationMiddleware(ctx *gin.Context) {
	authHeader := presenter.AuthHeaderStruct{}
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

	token := strings.TrimPrefix(authHeader.Authorization, "Bearer ")
	token = strings.TrimSpace(token)

	authHelper := helper.NewAuthBlueprint()
	claims, err := authHelper.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "token is invalid"})
		ctx.Abort()
		return
	}

	user := global.UsersRepo.GetUserByPhone(claims.Phone)
	claims, err = authHelper.ParseAndValidateToken(token, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "token is invalid"})
		ctx.Abort()
		return
	}

	ctx = SetClaimsToContext(ctx, claims)

	ctx.Next()
}

func SetClaimsToContext(ctx *gin.Context, claims presenter.PrivateClaims) *gin.Context {
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys[controller.ClaimsContextKey] = claims

	return ctx
}
