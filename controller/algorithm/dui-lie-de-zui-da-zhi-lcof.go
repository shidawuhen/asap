package algorithm

import "container/list"

/*
原题：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/
剑指 Offer 59 - II. 队列的最大值
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

//分析：
构造非严格递减函数。因为是从队列头部去掉，所以凡是比放入尾部的值小的，都可以删掉
*/

type MaxQueue struct {
	reallist *list.List
	maxlist  *list.List
}

func ConstructorMaxQueue() MaxQueue {
	return MaxQueue{
		reallist: list.New(),
		maxlist:  list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	e := this.maxlist.Front()
	if e == nil {
		return -1
	}
	return e.Value.(int)
}

func (this *MaxQueue) Push_back(value int) {
	this.reallist.PushBack(value)
	e := this.maxlist.Back()
	for e != nil && e.Value.(int) < value {
		this.maxlist.Remove(e)
		e = this.maxlist.Back()
	}
	this.maxlist.PushBack(value)

}

func (this *MaxQueue) Pop_front() int {
	e := this.reallist.Front()
	if e == nil {
		return -1
	}
	maxe := this.maxlist.Front()
	if maxe != nil && maxe.Value.(int) == e.Value.(int) {
		this.maxlist.Remove(maxe)
	}
	v := e.Value.(int)
	this.reallist.Remove(e)
	return v
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
