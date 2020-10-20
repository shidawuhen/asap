package algorithm

/*
原题：https://leetcode-cn.com/problems/reorder-list/
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

1234  1 4 3
*/

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	record := make([]*ListNode, 0)
	for head != nil {
		record = append(record, head)
		head = head.Next
	}
	start := record[0]
	head = record[0]
	length := len(record)

	for i := 0; i < length/2; i++ {
		start.Next = record[length-1-i]
		record[length-1-i].Next = record[i+1]
		start = record[i+1]
	}
	if length%2 != 0 {
		start.Next = record[length/2]
		start = record[length/2]
	}
	start.Next = nil
}
