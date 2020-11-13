package algorithm

/*
原题：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
剑指 Offer 25. 合并两个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	head := l1
	for l1.Next != nil && l2 != nil {
		//fmt.Print("start",l1.Val,l2.Val)
		//fmt.Printf("%v",l1)
		if l1.Next.Val >= l2.Val {
			l1Next := l1.Next
			l2Next := l2.Next
			l1.Next = l2
			l2.Next = l1Next
			l2 = l2Next
			l1 = l1.Next
		} else {
			l1 = l1.Next
		}
	}
	if l2 != nil {
		l1.Next = l2
	}
	return head
}
