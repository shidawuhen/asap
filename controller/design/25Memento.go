/**
@author: Jason Pang
@desc:
@date: 2021/8/15
**/
package design

import (
	"container/list"
	"fmt"
)

/**
 * @Author: Jason Pang
 * @Description: 备忘录
 */
type Memento struct {
	mario *Mario
}

func (m *Memento) GetMario() *Mario {
	return m.mario
}

/**
 * @Author: Jason Pang
 * @Description: 管理备忘录
 */
type Caretaker struct {
	stack *list.List
}

/**
 * @Author: Jason Pang
 * @Description: 保存备忘录
 * @receiver c
 * @param m
 */
func (c *Caretaker) Save(m *Memento) {
	c.stack.PushBack(m)
}

/**
 * @Author: Jason Pang
 * @Description: 获取上一个备忘录
 * @receiver c
 * @return *Memento
 */
func (c *Caretaker) Pop() *Memento {
	e := c.stack.Back()
	c.stack.Remove(e)
	return e.Value.(*Memento)
}

func mementoMain() {
	caretaker := &Caretaker{
		stack: list.New(),
	}

	mario := Mario{
		status: &SmallMarioStatus{},
		score:  0,
	}
	mario.status.SetMario(&mario)

	mario.status.Name()
	fmt.Println("-------------------获得蘑菇\n")
	mario.status.ObtainMushroom()

	mario.status.Name()
	fmt.Println("-------------------获得斗篷\n")
	mario.status.ObtainCape()

	fmt.Println("-------------------备份一下，要打怪了，当前状态为\n")
	mario.ShowInfo()
	caretaker.Save(mario.CreateMemento())
	fmt.Println("-------------------开始打怪\n")

	mario.status.Name()
	fmt.Println("-------------------遇到怪兽\n")
	mario.status.MeetMonster()

	fmt.Println("-------------------打怪失败，目前状态为\n")
	mario.ShowInfo()

	fmt.Println("-------------------恢复状态，重新打怪\n")
	mario.SetMemento(caretaker.Pop())
	mario.ShowInfo()
}
