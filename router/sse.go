/*
*
@author: Jason Pang
@desc:
@date: 2024/3/24
*
*/
package router

import (
	"asap/controller/sse"
	"github.com/gin-gonic/gin"
)

func sseFunc(router *gin.Engine) {
	router.GET("/sse/home", sse.Home())
	router.GET("/sse/subscribe", sse.Subscribe())
}
