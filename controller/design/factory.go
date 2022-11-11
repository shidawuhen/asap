/**
@author: Jason Pang
@desc:
@date: 2022/11/10
**/
package design

import "fmt"

/**
 * @Author: Jason Pang
 * @Description:
 */
type Life interface {
	CheckParams() error //参数检查
	Do() error          //执行动作
}

/**
 * @Author: Jason Pang
 * @Description: 基类
 */
type BaseLife struct {
}

/**
 * @Author: Jason Pang
 * @Description: 参数检查
 * @receiver c
 * @return error
 */
func (c *BaseLife) CheckParams() error {
	fmt.Println("通用参数检查")
	return nil
}

/**
 * @Author: Jason Pang
 * @Description: 开始做饭
 * @receiver c
 * @return error
 */
func (c *BaseLife) Do() error {
	fmt.Println("用水处理")
	return nil
}

type Cook struct {
}

/**
 * @Author: Jason Pang
 * @Description: 做饭参数检查
 * @receiver c
 * @return error
 */
func (c *Cook) CheckParams() error {
	fmt.Println("cook 检查参数，食材准备完毕")
	return nil
}

/**
 * @Author: Jason Pang
 * @Description: 开始做饭
 * @receiver c
 * @return error
 */
func (c *Cook) Do() error {
	fmt.Println("开始做饭")
	return nil
}

type Eat struct {
}

/**
 * @Author: Jason Pang
 * @Description: 吃饭参数检查
 * @receiver c
 * @return error
 */
func (c *Eat) CheckParams() error {
	fmt.Println("eat 检查参数，饭已做好，碗筷放好")
	return nil
}

/**
 * @Author: Jason Pang
 * @Description: 开始吃饭
 * @receiver c
 * @return error
 */
func (c *Eat) Do() error {
	fmt.Println("开始吃饭")
	return nil
}

/**
 * @Author: Jason Pang
 * @Description: 洗碗
 */
type Wash struct {
	BaseLife
}

type Mop struct {
	BaseLife
}

func (c *Mop) CheckParams() error {
	fmt.Println("mop 检查参数，拖把是否存在")
	return nil
}

/**
 * @Description: 简单工厂
 */
type Factory2 struct {
}

func (simple *Factory2) create(ext string) Life {
	switch ext {
	case "cook":
		return &Cook{}
	case "eat":
		return &Eat{}
	case "wash":
		return &Wash{}
	case "mop":
		return &Mop{}
	}
	return nil
}

func EchoBeforeDo() {
	fmt.Println("不想工作")
}

func factorymain() {
	//简单工厂使用代码
	fmt.Println("------------简单工厂")
	factory := &Factory2{}
	life := factory.create("mop")
	if life != nil {
		life.CheckParams()
		EchoBeforeDo()
		life.Do()
	}
}
