package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Activate(port int) {
	router := gin.Default()

	RegisterAuthAPI(router)

	portStr := ":" + strconv.Itoa(port)
	router.Run(portStr)
}
