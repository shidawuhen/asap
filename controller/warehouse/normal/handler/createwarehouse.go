/**
@author: Jason Pang
@desc:
@date: 2022/7/3
**/
package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatWareHouse(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
