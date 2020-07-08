package main

import (
	_ "asap/docs"
	"asap/router"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	router.InitRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}
