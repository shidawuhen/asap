/**
@author: Jason Pang
@desc:
@date: 2021/9/1
**/
package design

import "fmt"

/**
 * @Author: Jason Pang
 * @Description: 内容信息
 */
type Context struct {
	action  string
	content string
}

/**
 * @Author: Jason Pang
 * @Description: 翻译接口
 */
type Interpreter interface {
	Interpret(c Context)
}

/**
 * @Author: Jason Pang
 * @Description: 翻译音乐
 */
type MusicInterpreter struct {
}

/**
 * @Author: Jason Pang
 * @Description: 翻译音乐内容
 * @receiver m
 * @param c
 */
func (m MusicInterpreter) Interpret(c Context) {
	fmt.Println(c.action + " 中 " + c.content + " 的意思是感情高昂")
}

/**
 * @Author: Jason Pang
 * @Description: 翻译舞蹈
 */
type DanceInterpreter struct {
}

/**
 * @Author: Jason Pang
 * @Description: 翻译舞蹈内容
 * @receiver d
 * @param c
 */
func (d DanceInterpreter) Interpret(c Context) {
	fmt.Println(c.action + " 中 " + c.content + " 的意思是悲凉")
}

func interpreterMain() {
	cList := []Context{
		{action: "music", content: "高音"},
		{action: "music", content: "低音"},
		{action: "dance", content: "跳跃"},
		{action: "dance", content: "挥手"},
	}
	//对歌舞剧内容进行翻译
	for _, c := range cList {
		if c.action == "music" {
			MusicInterpreter{}.Interpret(c)
		} else if c.action == "dance" {
			DanceInterpreter{}.Interpret(c)
		}
	}
}
