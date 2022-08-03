/**
@author: Jason Pang
@desc:
@date: 2022/8/3
**/
package list

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @Description:
	比较回文的思路，是一个从头部，一个从尾部开始对两个值做对比，比较到中点处结束
	1. 写到数组里，然后在数组里比较
	2. 改造成双向链表，一个往前一个往后
	3. 后半部分改成从后往前的。
    需要找到中间位置
    - 方案1：遍历一遍，记录总长度，再遍历一次，就知道中间位置在哪里了
    - 方案2：快慢指针，快指针走到结束，慢指针走到中间。其实快慢指针是我们1/2,1/3等思想的具体实践
要走都走，无论奇数偶数，慢指针走到的是最后一个
	然后从中间位置开始改变链表顺序
	一个从头，一个从尾开始遍历，从尾部走，走到空就行


*/
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	//找到中间位置
	slow := head
	fast := head
	listLen := 0
	for fast != nil && fast.Next != nil {
		if listLen%2 == 0 {
			slow = slow.Next
		}
		fast = fast.Next
		listLen++
	}
	//从slow到fast做反向
	tmp := slow.Next
	slow.Next = nil
	var end *ListNode = nil
	if tmp != nil {
		end = tmp.Next
	}

	for tmp != nil {
		tmp.Next = slow
		slow = tmp
		tmp = end
		if tmp != nil {
			end = tmp.Next
		}
	}
	//从头到中间、从尾到中间
	h := head
	for fast != nil {
		if fast.Val != h.Val {
			return false
		}
		fast = fast.Next
		h = h.Next
	}
	return true
}
