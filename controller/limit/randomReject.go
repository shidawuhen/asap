package limit

import (
"github.com/gin-gonic/gin"
"math/rand"
"net/http"
)

// @Tags limit
// @Summary 随机拒流
// @Produce  json
// @Success 200 {string} string "成功会返回ok"
// @Failure 502 "失败返回reject"
// @Router /limit/randomreject [get]
func RandomReject(c *gin.Context) {
	refuseRate := 100
	if refuseRate != 0 {
		temp := rand.Intn(1000)
		if temp <= refuseRate {
			c.String(http.StatusBadGateway, "reject")
			return
		}
	}
	c.String(http.StatusOK, "ok")
}
