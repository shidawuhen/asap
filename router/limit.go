package router

import (
	"asap/controller/limit"
	"github.com/gin-gonic/gin"
)

func limitFunc(router *gin.Engine) {
	router.GET("/limit/randomreject",  limit.RandomReject)
}
