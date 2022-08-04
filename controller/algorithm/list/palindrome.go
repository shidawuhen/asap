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
	//从slow到fast做反向。不应该用这种方案，为啥要next、nn呢，要是用三个的话，当前的cur，然后找前一个和后一个，把自己放到中间最好处理
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

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	//找到中间位置，这种方案需要slow往后走一步
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//从slow往后逆序
	cur := slow.Next
	var pre *ListNode = nil
	for cur != nil {
		tmpNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmpNext
	}
	//从头到中间、从尾到中间
	h := head
	for pre != nil {
		if pre.Val != h.Val {
			return false
		}
		pre = pre.Next
		h = h.Next
	}
	return true
}
