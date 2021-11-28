/**
@author: Jason Pang
@desc: http状态码
@date: 2021/11/28
**/
package various

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Code200(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func Code500(c *gin.Context) {
	panic("1")
	c.String(http.StatusOK, "ok")
}

func Code502(c *gin.Context) {
	//方案一
	/*
		package main

		import (
			"fmt"
			"net"
		)

		func main() {
			ln, err := net.Listen("tcp", "127.0.0.1:8082")
			if err != nil {
				return
			}
			go func() {
				for {
					c, err := ln.Accept()
					fmt.Println("Accept")
					if err != nil {
						break
					}
					c.Close()
				}
			}()
			select {}
		}
	*/
	//方案二
	/*
		//r.Run(":8082")
		server := http.Server{
		   Addr:         ":8082",
		   WriteTimeout: time.Second * 1,
		   ReadTimeout:  time.Second * 10,
		   IdleTimeout:  time.Second * 10,
		   Handler:      r,
		}

		server.ListenAndServe()
	*/
	time.Sleep(time.Second * 2)
	c.String(http.StatusOK, "ok")
}

func Code504(c *gin.Context) {
	time.Sleep(time.Second * 100)
	c.String(http.StatusOK, "ok")
}

func Code503(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func Code499(c *gin.Context) {
	time.Sleep(time.Second * 100)
	c.String(http.StatusOK, "ok")
}
