/*
*
@author: Jason Pang
@desc:
@date: 2024/3/24
*
*/
package sse

import (
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func Subscribe() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Stream message to client
		count := 0
		c.Stream(func(w io.Writer) bool {
			timer := time.NewTicker(time.Second)
			for range timer.C {
				c.SSEvent("message", "hello "+time.Now().String())
				count++
				if count > 10 {
					return false
				}
				return true
			}
			return false
		})
	}
}
