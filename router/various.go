/**
@author: Jason Pang
@desc:
@date: 2021/11/28
**/
package router

import (
	"asap/controller/various"
	"github.com/gin-gonic/gin"
)

func variousFunc(router *gin.Engine) {
	router.GET("/httpcode/code200", various.Code200)
	router.GET("/httpcode/code500", various.Code500)
	router.GET("/httpcode/code502", various.Code502)
	router.GET("/httpcode/code504", various.Code504)
	router.GET("/httpcode/code503", various.Code503)
	router.GET("/httpcode/code499", various.Code499)
}
