/**
@date: 2021/5/24
**/
package router

import (
	"asap/controller/paramtype"
	"github.com/gin-gonic/gin"
)

func paramTypeFunc(router *gin.Engine) {
	router.GET("/paramtype/query/:param", paramtype.Query)
	router.POST("/paramtype/postformdata", paramtype.PostFormData)
	router.POST("/paramtype/raw/:name", paramtype.Raw)
	router.POST("/paramtype/binary", paramtype.Binary)
}
