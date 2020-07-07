package main

import (
	"net/http"

	_ "asap/docs"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", ping)

	return r
}

// @Summary 接口探活
// @Produce  json
// @Param lang query string false "en"
// @Success 200 {string} string "ok"
// @Router /ping [get]
func ping(c *gin.Context) {
	cookies := c.Request.Cookies()
	cookieInfo := ""
	for _, cookie := range cookies{
		cookieInfo += cookie.Name + ":" + cookie.Value + "\n"
	}
	c.String(http.StatusOK, cookieInfo )

}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9090")
}
