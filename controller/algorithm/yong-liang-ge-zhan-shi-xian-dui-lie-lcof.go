package algorithm

import "strings"

/*
原题：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用

分析：
多玩了一下，每次操作后都将real置为正确的
*/

type CQueue struct {
	real   []int
	backup []int
}

func Constructor() CQueue {
	return CQueue{
		real:   make([]int, 0),
		backup: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	//将real的数据放入到backup，然后将value放入backup，再将backup放回到real
	this.backup = make([]int, 0)
	for i := len(this.real) - 1; i >= 0; i-- {
		this.backup = append(this.backup, this.real[i])
	}
	this.real = make([]int, 0)
	this.backup = append(this.backup, value)
	for i := len(this.backup) - 1; i >= 0; i-- {
		this.real = append(this.real, this.backup[i])
	}
	strings.TrimLeft()
	strings.TrimLeft("abc", "@#$!%^&*()_+=-"))
}

func (this *CQueue) DeleteHead() int {
	if len(this.real) == 0 {
		return -1
	}
	if len(this.real) == 1 {
		d := this.real[0]
		this.real = make([]int, 0)
		return d
	}
	d := this.real[len(this.real)-1]
	this.real = this.real[0 : len(this.real)-1]
	return d
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
