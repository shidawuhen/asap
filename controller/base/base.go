package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 接口探活
// @Produce  json
// @Param lang query string false "en"
// @Success 200 {string} string "ok"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
