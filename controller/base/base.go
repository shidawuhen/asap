package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags base
// @Summary 接口探活
// @Description ping接口，用于测试是否调通
// @Produce  json
// @Param name query string false "参数描述"
// @Success 200 {string} string "ok"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
