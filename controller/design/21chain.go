/**
@author: Jason Pang
@desc:
@date: 2021/7/21
**/
package design

import "fmt"

var status int8 = 0

type HandlerFunc func()

type HandlersChain []HandlerFunc

/**
 * @Author: Jason Pang
 * @Description:
 */
type RouterGroup struct {
	Handlers HandlersChain
	index    int8
}

/**
 * @Author: Jason Pang
 * @Description: 添加中间件，将其组成链式
 * @receiver group
 * @param middleware
 */
func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
}

/**
 * @Author: Jason Pang
 * @Description: 链顺序执行
 * @receiver group
 */
func (group *RouterGroup) Next() {
	for group.index < int8(len(group.Handlers)) {
		group.Handlers[group.index]()
		group.index++
	}
}

/**
 * @Author: Jason Pang
 * @Description: 中间件
 */
func middleware1() {
	fmt.Println("全局中间件1执行完毕")
}

/**
 * @Author: Jason Pang
 * @Description: 中间件
 */
func middleware2() {
	fmt.Println("全局中间件2执行失败")
	status = 1
}

func chainMain() {
	r := &RouterGroup{}
	//添加中间件
	r.Use(middleware1, middleware2)
	//运行中间件
	r.Next()
	//状态检查
	if status == 1 {
		fmt.Println("中间件检查失败，请重试")
		return
	}
	//执行后续流程
}
