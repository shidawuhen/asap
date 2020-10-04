package algorithm

/*
原题：https://leetcode-cn.com/problems/add-two-numbers/
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807


*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, index *ListNode = nil, nil
	another := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + another
		another = val / 10
		val = val % 10
		if head == nil {
			head = &ListNode{Val: val}
			index = head
		} else {
			index.Next = &ListNode{Val: val}
			index = index.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := l1.Val + another
		another = val / 10
		val = val % 10
		if head == nil {
			head = &ListNode{Val: val}
			index = head
		} else {
			index.Next = &ListNode{Val: val}
			index = index.Next
		}
		l1 = l1.Next
	}
	for l2 != nil {
		val := l2.Val + another
		another = val / 10
		val = val % 10
		if head == nil {
			head = &ListNode{Val: val}
			index = head
		} else {
			index.Next = &ListNode{Val: val}
			index = index.Next
		}
		l2 = l2.Next
	}
	if another != 0 {
		index.Next = &ListNode{Val: another}
		index = index.Next
	}
	return head
}
