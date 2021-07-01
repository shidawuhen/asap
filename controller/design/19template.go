/**
@author: Jason Pang
@desc:
@date: 2021/7/1
**/
package design

import "fmt"

/**
 * @Author: Jason Pang
 * @Description: 试卷
 */
type Examination struct {
	//函数变量，回答问题1
	Answer1 func()
	//函数变量，回答问题2
	Answer2 func()
}

/**
 * @Author: Jason Pang
 * @Description: 问题列表，也是算法骨架
 * @receiver e
 */
func (e *Examination) Questions() {
	fmt.Println("第一题：谁是最帅的人？")
	e.Answer1()
	fmt.Println("第二题：生活的意义是什么？")
	e.Answer2()
}

/**
 * @Author: Jason Pang
 * @Description: 真正做试卷
 */
type ExamplationDo struct {
	Examination
}

/**
 * @Author: Jason Pang
 * @Description: 写答案1
 * @receiver d
 */
func (d *ExamplationDo) Answer1() {
	fmt.Println("答案：我自己")
}

/**
 * @Author: Jason Pang
 * @Description: 写答案2
 * @receiver d
 */
func (d *ExamplationDo) Answer2() {
	fmt.Println("答案：躺平")
}
func templateMain() {
	e := &ExamplationDo{}
	//需要对父类函数进行赋值
	e.Examination.Answer1 = e.Answer1
	e.Examination.Answer2 = e.Answer2

	e.Questions()
}
