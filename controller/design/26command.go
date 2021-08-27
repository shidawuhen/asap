/**
@author: Jason Pang
@desc:
@date: 2021/8/27
**/
package design

import "fmt"

/**
 * @Author: Jason Pang
 * @Description: 命令接口
 */
type Command interface {
	Execute()
}

/**
 * @Author: Jason Pang
 * @Description: 移动命令
 */
type MoveCommand struct {
	x, y int64
}

/**
 * @Author: Jason Pang
 * @Description: 如何移动
 * @receiver m
 */
func (m *MoveCommand) Execute() {
	fmt.Printf("向右移动%d，向上移动%d \n", m.x, m.y)
}

/**
 * @Author: Jason Pang
 * @Description: 攻击命令
 */
type AttackCommand struct {
	skill string
}

/**
 * @Author: Jason Pang
 * @Description: 如何攻击
 * @receiver a
 */
func (a *AttackCommand) Execute() {
	fmt.Printf("使用技能%s\n", a.skill)
}

/**
 * @Author: Jason Pang
 * @Description: 记录命令
 * @param action
 * @return Command
 */
func AddCommand(action string) Command {
	if action == "attack" {
		return &AttackCommand{
			skill: "野蛮冲撞",
		}
	} else { //默认是移动
		return &MoveCommand{
			x: 10,
			y: 20,
		}
	}
}

func commandMain() {
	//将命令记录
	lc := make([]Command, 0)
	lc = append(lc, AddCommand("attack"))
	lc = append(lc, AddCommand("move"))
	lc = append(lc, AddCommand("move"))
	lc = append(lc, AddCommand("attack"))

	//执行命令
	for _, c := range lc {
		c.Execute()
	}
}
