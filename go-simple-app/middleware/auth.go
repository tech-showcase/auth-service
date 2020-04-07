package middleware

import (
	"github.com/gin-gonic/gin"
	"go-simple-app/controller"
	"go-simple-app/helper"
	"go-simple-app/presenter"
	"go-simple-app/singleton"
	"net/http"
	"strings"
)

func JWTAuthenticationMiddleware(ctx *gin.Context) {
	authHeader := presenter.AuthHeaderStruct{}
	if err := ctx.ShouldBindHeader(&authHeader); err != nil || authHeader.Authorization == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "request should contains authorization header"})
		return
	}

	if !strings.Contains(authHeader.Authorization, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "token should contains Bearer"})
		return
	}

	token := strings.TrimPrefix(authHeader.Authorization, "Bearer ")
	token = strings.TrimSpace(token)

	authHelper := helper.NewAuthBlueprint()
	claims, err := authHelper.ParseTokenWithoutKey(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "token is invalid"})
		return
	}

	user := singleton.UsersRepo.GetUserByPhone(claims.Phone)
	claims, err = authHelper.ParseToken(token, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "token is invalid"})
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
