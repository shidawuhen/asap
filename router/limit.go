package router

import (
	"asap/controller/limit"
	"github.com/gin-gonic/gin"
)

func limitFunc(router *gin.Engine) {
	router.GET("/limit/randomreject",  limit.RandomReject)
	router.GET("/limit/countreject",  limit.CountReject)
	router.GET("/limit/slidewindowsreject",  limit.SlideWindowsReject)
}






