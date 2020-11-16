package algorithm

/*
原题：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
剑指 Offer 30. 包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。



示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.


提示：

各函数的调用总次数不超过 20000 次

分析：
好好做题，不要自己瞎增加难度
*/
type MinStack struct {
	stack           []int
	stackIndex      int
	lowerStack      []int
	lowerStackIndex int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:           make([]int, 0),
		stackIndex:      -1,
		lowerStack:      make([]int, 0),
		lowerStackIndex: -1,
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == this.stackIndex+1 { //空间不够了
		this.stack = append(this.stack, x)
	} else {
		this.stack[this.stackIndex+1] = x
	}
	this.stackIndex++
	if this.lowerStackIndex == -1 || this.lowerStack[this.lowerStackIndex] >= x {
		if len(this.lowerStack) == this.lowerStackIndex+1 {
			this.lowerStack = append(this.lowerStack, x)
		} else {
			this.lowerStack[this.lowerStackIndex+1] = x
		}
		this.lowerStackIndex++
	}
	//fmt.Println(this.stack,this.stackIndex,this.lowerStack,this.lowerStackIndex)
}

func (this *MinStack) Pop() {
	if this.stackIndex != -1 {
		v := this.stack[this.stackIndex]
		if this.lowerStackIndex != -1 && v == this.lowerStack[this.lowerStackIndex] {
			this.lowerStackIndex--
		}
		this.stackIndex--
	}
}

func (this *MinStack) Top() int {
	if this.stackIndex == -1 {
		return 0
	}
	return this.stack[this.stackIndex]
}

func (this *MinStack) Min() int {
	if this.lowerStackIndex == -1 {
		return 0
	}
	return this.lowerStack[this.lowerStackIndex]
}
