package router

import (
	"asap/controller/grpcclient"
	"github.com/gin-gonic/gin"
)

func grpcFunc(router *gin.Engine) {
	router.GET("/grpcclient/hello",  grpcclient.Hello)
}
