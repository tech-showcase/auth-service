package controller

import (
	"github.com/gin-gonic/gin"
	"go-simple-app/helper"
	"go-simple-app/presenter"
	"net/http"
)

func Register(ctx *gin.Context) {
	registerRequest := presenter.RegisterRequestStruct{}
	if errA := ctx.ShouldBind(&registerRequest); errA == nil {
		password := helper.Generate4CharsPassword(registerRequest.Phone)
		registerResponse := presenter.RegisterResponseStruct{
			Password: password,
		}

		ctx.JSON(http.StatusOK, registerResponse)
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "request body should contains JSON"})
	}
}

func Login(ctx *gin.Context) {

}
