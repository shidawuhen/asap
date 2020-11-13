package algorithm

/*
原题：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
剑指 Offer 24. 反转链表
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000

*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head
	pro := head.Next
	newHead.Next = nil
	for pro != nil {
		temp := pro.Next
		pro.Next = newHead
		newHead = pro
		pro = temp
	}
	return newHead
}
