/**
@date: 2021/5/25
**/
package paramtype

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostFormStruct struct {
	GetPostForm string `json:"getPostForm" uri:"getPostForm" form:"getPostForm"`
}

/**
 * @Author: POST form数据获取
 * @Description:
 * @param c
 */
func PostFormData(c *gin.Context) {
	//需定义合规结构体
	postFormStruct := &PostFormStruct{}
	c.ShouldBind(postFormStruct)
	fmt.Printf("%+v \n", postFormStruct)

	postForm := c.PostForm("postForm")
	fmt.Println(postForm)

	defaultPostForm := c.DefaultPostForm("defaultPostForm", "no")
	fmt.Println(defaultPostForm)

	getPostForm, res := c.GetPostForm("getPostForm")
	fmt.Println(getPostForm, res)

	postFormArray := c.PostFormArray("postFormArray")
	fmt.Println(postFormArray)

	getPostFormArray, res := c.GetPostFormArray("getPostFormArray")
	fmt.Println(getPostFormArray, res)

	postFormMap := c.PostFormMap("postFormMap")
	fmt.Println(postFormMap)

	getPostFormMap, res := c.GetPostFormMap("getPostFormMap")
	fmt.Println(getPostFormMap)

	c.String(http.StatusOK, "ok")
}
