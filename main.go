package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// Ping test
	r.GET("/ping", ping)
	r.GET("/longconnecthtml",longconnecthtml)

	return r
}

func longconnecthtml(c *gin.Context)  {
	c.HTML(http.StatusOK, "longconnect.tmpl",gin.H{})
}

// @Summary 接口探活
// @Produce  json
// @Param lang query string false "en"
// @Success 200 {string} string "ok"
// @Router /ping [get]
func ping(c *gin.Context) {
	//c.String(http.StatusOK, "ok")
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}

	}
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9090")
}
