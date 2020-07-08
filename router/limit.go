package router

import (
	"asap/controller/base"
	"github.com/gin-gonic/gin"
)

func limit(router *gin.Engine) {
	router.GET("/limit/ping",  base.Ping)
}
