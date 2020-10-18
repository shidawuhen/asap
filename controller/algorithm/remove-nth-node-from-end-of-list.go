package algorithm

/*
原题：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first, second := head, head
	if head != nil {
		for n > 0 {
			second = second.Next
			n--

		}
		for second != nil && second.Next != nil {
			second = second.Next
			first = first.Next
		}

		if second == nil {
			head = head.Next
		} else {
			first.Next = first.Next.Next
		}

	}
	return head
}
