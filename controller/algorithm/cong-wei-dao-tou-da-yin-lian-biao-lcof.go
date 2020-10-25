package algorithm

/*
原题：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
剑指 Offer 06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]


限制：

0 <= 链表长度 <= 10000

*/

func reversePrint(head *ListNode) []int {
	record := make([]int, 0)
	for head != nil {
		record = append(record, head.Val)
		head = head.Next
	}
	length := len(record)
	for i := 0; i < length/2; i++ {
		record[i], record[length-1-i] = record[length-1-i], record[i]
	}
	return record
}
