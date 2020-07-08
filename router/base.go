package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"asap/controller/base"
)

func baseFunc(router *gin.Engine) {
	// Ping test
	router.GET("/ping", base.Ping)

	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

