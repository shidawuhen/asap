/**
@date: 2021/5/26
**/
package paramtype

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description: 获取Binary数据，没有发现指定的函数
 * @param c
 */
func Binary(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
