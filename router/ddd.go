/**
@author: Jason Pang
@desc:
@date: 2022/7/4
**/
package router

import (
	"asap/controller/warehouse/normal/handler"
	"github.com/gin-gonic/gin"
)

func dddFunc(router *gin.Engine) {
	router.GET("/nddd/getwarehouse", handler.GetWareHouse)
	router.GET("/nddd/createwarehouse", handler.CreatWareHouse)
}
