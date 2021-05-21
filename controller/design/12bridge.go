/**
@date: 2021/5/21
**/
package design

import "fmt"

/**
 * @Description: 消息发送接口
 */
type MessageSend interface {
	send(msg string)
}

/**
 * @Description: 短信消息
 */
type SMSSend struct {
}

func (s *SMSSend) send(msg string) {
	fmt.Println("sms 发送的消息内容为: " + msg)
}

/**
 * @Description: 邮件消息
 */
type Email struct {
}

func (e *Email) send(msg string) {
	fmt.Println("email 发送的消息内容为: " + msg)
}

/**
 * @Description: AppPush消息
 */
type AppPush struct {
}

func (a *AppPush) send(msg string) {
	fmt.Println("appPush 发送的消息内容为: " + msg)
}

/**
 * @Description: 站内信消息
 */
type Letter struct {
}

func (l *Letter) send(msg string) {
	fmt.Println("站内信 发送的消息内容为: " + msg)
}

/**
 * @Description: 用户触达父类，包含触达方式数组messageSends
 */
type Touch struct {
	messageSends []MessageSend
}

/**
 * @Description: 触达方法，调用每一种方式进行触达
 * @receiver t
 * @param msg
 */
func (t *Touch) do(msg string) {
	for _, s := range t.messageSends {
		s.send(msg)
	}
}

/**
 * @Description: 紧急消息做用户触达
 */
type TouchUrgent struct {
	base Touch
}

/**
 * @Description: 紧急消息，先从db中获取各种信息，然后使用各种触达方式通知用户
 * @receiver t
 * @param msg
 */
func (t *TouchUrgent) do(msg string) {
	fmt.Println("touch urgent 从db获取接收人等信息")
	t.base.do(msg)
}

/**
 * @Description: 普通消息做用户触达
 */
type TouchNormal struct {
	base Touch
}

/**
 * @Description: 普通消息，先从文件中获取各种信息，然后使用各种触达方式通知用户
 * @receiver t
 * @param msg
 */
func (t *TouchNormal) do(msg string) {
	fmt.Println("touch normal 从文件获取接收人等信息")
	t.base.do(msg)
}

func bridgeMain() {
	//触达方式
	sms := &SMSSend{}
	appPush := &AppPush{}
	letter := &Letter{}
	email := &Email{}
	//根据触达类型选择触达方式
	fmt.Println("-------------------touch urgent")
	touchUrgent := TouchUrgent{
		base: Touch{
			messageSends: []MessageSend{sms, appPush, letter, email},
		},
	}
	touchUrgent.do("urgent情况")
	fmt.Println("-------------------touch normal")
	touchNormal := TouchNormal{ //
		base: Touch{
			messageSends: []MessageSend{sms, appPush, letter, email},
		},
	}
	touchNormal.do("normal情况")
}
