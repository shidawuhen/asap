/**
@author: Jason Pang
@desc:
@date: 2021/9/4
**/
package design

import "fmt"

/**
 * @Author: Jason Pang
 * @Description: 中介者接口
 */
type Mediator interface {
	Process(colleague UIColleague)
}

/**
 * @Author: Jason Pang
 * @Description: 真正的中介者
 */
type UIMediator struct {
	activate *ActivateColleague
	evaluate *EvaluateColleague
	text     *TextColleague
}

/**
 * @Author: Jason Pang
 * @Description: 中介者大总管，处理所有事情
 * @receiver u
 */
func (u *UIMediator) Process(colleague UIColleague) {
	if colleague == u.activate { //如果是激活
		u.evaluate.Show("试用内容隐藏")
		u.text.Show("请输入激活码")
	} else if colleague == u.evaluate { //如果是试用
		u.activate.Show("激活内容隐藏")
		u.text.Show("请出入激活时间")
	}
}

/**
 * @Author: Jason Pang
 * @Description: Colleague接口
 */
type UIColleague interface {
	Action()
}

/**
 * @Author: Jason Pang
 * @Description: 激活UI
 */
type ActivateColleague struct {
	mediator Mediator
}

/**
 * @Author: Jason Pang
 * @Description: 激活触发的动作
 * @receiver a
 */
func (a *ActivateColleague) Action() {
	a.mediator.Process(a)
}

/**
 * @Author: Jason Pang
 * @Description: 激活UI显示内容
 * @receiver e
 * @param text
 */
func (e *ActivateColleague) Show(text string) {
	fmt.Println(text)
}

/**
 * @Author: Jason Pang
 * @Description: 试用UI
 */
type EvaluateColleague struct {
	mediator Mediator
}

/**
 * @Author: Jason Pang
 * @Description: 试用触发的动作
 * @receiver e
 */
func (e *EvaluateColleague) Action() {
	e.mediator.Process(e)
}

/**
 * @Author: Jason Pang
 * @Description: 试用UI显示内容
 * @receiver e
 * @param text
 */
func (e *EvaluateColleague) Show(text string) {
	fmt.Println(text)
}

/**
 * @Author: Jason Pang
 * @Description: 文案UI
 */
type TextColleague struct {
	mediator Mediator
}

/**
 * @Author: Jason Pang
 * @Description: 文案触发动作
 * @receiver t
 */
func (t *TextColleague) Action() {
	t.mediator.Process(t)
}

/**
 * @Author: Jason Pang
 * @Description: 文案显示内容
 * @receiver t
 * @param text
 */
func (t *TextColleague) Show(text string) {
	fmt.Println(text)
}

func mediatorMain() {
	//初始化
	m := &UIMediator{}

	activate := &ActivateColleague{
		mediator: m,
	}
	evaluate := &EvaluateColleague{
		mediator: m,
	}
	text := &TextColleague{
		mediator: m,
	}

	m.activate = activate
	m.evaluate = evaluate
	m.text = text
	//点击激活
	fmt.Println("-----------------点击激活")
	activate.Action()
	//点击试用
	fmt.Println("-----------------点击试用")
	evaluate.Action()
}
