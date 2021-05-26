/**
@author: pangzhiqiang
@date: 2021/5/26
**/
package paramtype

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Binary(c *gin.Context) {

	//c.QueryMap()
	c.String(http.StatusOK, "ok")
}
