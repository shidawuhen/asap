/**
@date: 2021/5/26
**/
package paramtype

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RawStruct struct {
	Id   string
	Text string
}

type RawUri struct {
	Name string `json:"name" uri:"name"`
}

type RawQuery struct {
	Query string `json:"query" uri:"query" form:"query"`
}

/**
 * @Description: POST Raw数据获取
 * @param c
 */
func Raw(c *gin.Context) {

	raw := &RawStruct{}
	c.ShouldBind(raw)
	fmt.Printf("%+v \n", raw)

	//需定义合规结构体
	rawUri := &RawUri{}
	c.ShouldBindUri(rawUri)
	fmt.Printf("%+v \n", rawUri)

	//需定义合规结构体
	rawQuery := &RawQuery{}
	c.ShouldBindQuery(rawQuery)
	fmt.Printf("%+v \n", rawQuery)

	c.String(http.StatusOK, "ok")
}
