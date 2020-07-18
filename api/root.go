package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/config"
	"go.elastic.co/apm/module/apmgin"
	"io"
	"os"
	"strconv"
)

func Activate(port int) {
	setupHTTPLogger()

	router := gin.Default()
	setupAPMAgent(router)

	RegisterSupportAPI(router)
	RegisterJWTAPI(router)
	RegisterOAuth2API(router)

	address := ":" + strconv.Itoa(port)
	router.Run(address)
}

func setupHTTPLogger() {
	f, err := os.OpenFile(config.Instance.Log.FilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func setupAPMAgent(router *gin.Engine) {
	router.Use(apmgin.Middleware(router))
}
