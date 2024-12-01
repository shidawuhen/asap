package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	baseFunc(router)
	limitFunc(router)
	grpcFunc(router)
	paramTypeFunc(router)
	variousFunc(router)
	dddFunc(router)
	sseFunc(router)
	oauthFunc(router)
}
