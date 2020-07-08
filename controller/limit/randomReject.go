package limit

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

// @Tags limit
// @Summary 随机拒流
// @Produce  json
// @Param name query string false "参数描述"
// @Success 200 {string} string "ok"
// @Router /limit/randomreject [get]
func RandomReject(c *gin.Context) {
	refuseRate := 200
	if refuseRate != 0 {
		temp := rand.Intn(1000)
		if temp <= refuseRate {
			c.String(http.StatusOK, "reject")
			return
		}
	}
	c.String(http.StatusOK, "ok")
}
