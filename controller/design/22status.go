/**
@author: Jason Pang
@desc:
@date: 2021/8/3
**/
package design

import "fmt"

type Mario struct {
	score  int64
	status MarioStatus
}

/**
 * @Author: Jason Pang
 * @Description: 展示信息和分数
 * @receiver m
 */
func (m *Mario) ShowInfo() {
	m.status.Name()
	fmt.Println("当前分数为:", m.score)
}

/**
 * @Author: Jason Pang
 * @Description: 创建备忘录
 * @receiver m
 */
func (m *Mario) CreateMemento() *Memento {
	return &Memento{
		mario: &Mario{
			score:  m.score,
			status: m.status,
		},
	}
}

/**
 * @Author: Jason Pang
 * @Description: 恢复数据
 * @receiver m
 * @param mem
 */
func (m *Mario) SetMemento(mem *Memento) {
	m.score = mem.mario.score
	m.status = mem.mario.status
}

type MarioStatus interface {
	Name()
	ObtainMushroom()
	ObtainCape()
	MeetMonster()
	SetMario(mario *Mario)
}

/**
 * @Author: Jason Pang
 * @Description: 小马里奥
 */
type SmallMarioStatus struct {
	mario *Mario
}

/**
 * @Author: Jason Pang
 * @Description: 设置马里奥
 * @receiver s
 * @param mario
 */
func (s *SmallMarioStatus) SetMario(mario *Mario) {
	s.mario = mario
}

func (s *SmallMarioStatus) Name() {
	fmt.Println("小马里奥")
}

/**
 * @Author: Jason Pang
 * @Description: 获得蘑菇变为超级马里奥
 * @receiver s
 */
func (s *SmallMarioStatus) ObtainMushroom() {
	s.mario.status = &SuperMarioStatus{
		mario: s.mario,
	}
	s.mario.score += 100
}

/**
 * @Author: Jason Pang
 * @Description: 获得斗篷变为斗篷马里奥
 * @receiver s
 */
func (s *SmallMarioStatus) ObtainCape() {
	s.mario.status = &CapeMarioStatus{
		mario: s.mario,
	}
	s.mario.score += 200
}

/**
 * @Author: Jason Pang
 * @Description: 遇到怪兽减100
 * @receiver s
 */
func (s *SmallMarioStatus) MeetMonster() {
	s.mario.score -= 100
}

/**
 * @Author: Jason Pang
 * @Description: 超级马里奥
 */

type SuperMarioStatus struct {
	mario *Mario
}

/**
 * @Author: Jason Pang
 * @Description: 设置马里奥
 * @receiver s
 * @param mario
 */
func (s *SuperMarioStatus) SetMario(mario *Mario) {
	s.mario = mario
}

func (s *SuperMarioStatus) Name() {
	fmt.Println("超级马里奥")
}

/**
 * @Author: Jason Pang
 * @Description: 获得蘑菇无变化
 * @receiver s
 */
func (s *SuperMarioStatus) ObtainMushroom() {

}

/**
 * @Author: Jason Pang
 * @Description:获得斗篷变为斗篷马里奥
 * @receiver s
 */
func (s *SuperMarioStatus) ObtainCape() {
	s.mario.status = &CapeMarioStatus{
		mario: s.mario,
	}
	s.mario.score += 200
}

/**
 * @Author: Jason Pang
 * @Description: 遇到怪兽变为小马里奥
 * @receiver s
 */
func (s *SuperMarioStatus) MeetMonster() {
	s.mario.status = &SmallMarioStatus{
		mario: s.mario,
	}
	s.mario.score -= 200
}

/**
 * @Author: Jason Pang
 * @Description: 斗篷马里奥
 */
type CapeMarioStatus struct {
	mario *Mario
}

/**
 * @Author: Jason Pang
 * @Description: 设置马里奥
 * @receiver s
 * @param mario
 */
func (c *CapeMarioStatus) SetMario(mario *Mario) {
	c.mario = mario
}

func (c *CapeMarioStatus) Name() {
	fmt.Println("斗篷马里奥")
}

/**
 * @Author: Jason Pang
 * @Description:获得蘑菇无变化
 * @receiver c
 */
func (c *CapeMarioStatus) ObtainMushroom() {

}

/**
 * @Author: Jason Pang
 * @Description: 获得斗篷无变化
 * @receiver c
 */
func (c *CapeMarioStatus) ObtainCape() {

}

/**
 * @Author: Jason Pang
 * @Description: 遇到怪兽变为小马里奥
 * @receiver c
 */
func (c *CapeMarioStatus) MeetMonster() {
	c.mario.status = &SmallMarioStatus{
		mario: c.mario,
	}
	c.mario.score -= 200
}
func statusMain() {
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

	mario.status.Name()
	fmt.Println("-------------------遇到怪兽\n")
	mario.status.MeetMonster()

	mario.status.Name()
}
