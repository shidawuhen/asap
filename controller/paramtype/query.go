/**
@date: 2021/5/24
**/
package paramtype

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description: GET Query参数获取实例
 * @param c
 */
func Query(c *gin.Context) {

	param := c.Param("param")
	fmt.Println(param)

	query := c.Query("query")
	fmt.Println(query)

	defaultQuery := c.DefaultQuery("defaultQuery", "no")
	fmt.Println(defaultQuery)

	getQuery, res := c.GetQuery("getQuery")
	fmt.Println(getQuery, res)

	queryArray := c.QueryArray("queryArray")
	fmt.Println(queryArray)

	getQueryArray, res := c.GetQueryArray("getQueryArray")
	fmt.Println(getQueryArray, res)

	queryMap := c.QueryMap("queryMap")
	fmt.Println(queryMap)

	getQueryMap, res := c.GetQueryMap("getQueryMap")
	fmt.Println(getQueryMap)

	c.String(http.StatusOK, "ok")
}
