package main

import (
	"asap/router"
	"net/http"

	_ "asap/docs"

	"github.com/gin-gonic/gin"
)


// @Summary 接口探活
// @Produce  json
// @Param lang query string false "en"
// @Success 200 {string} string "ok"
// @Router /ping [get]
func ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func main() {
	r := gin.Default()
	router.InitRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}
