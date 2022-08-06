/**
@author: Jason Pang
@desc:
@date: 2022/8/4
**/
package list

/**
合并两个有序列表，这个不需要太多技巧，纯看能力
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//比较大小，确保l1肯定是最小的
	if l1.Val > l2.Val {
		tmp := l1
		l1 = l2
		l2 = tmp
	}
	l1Ptr := l1
	//从l1开始，如果l2比它小，就加进来
	for l1Ptr != nil && l1Ptr.Next != nil {
		if l2 != nil && l2.Val <= l1Ptr.Next.Val {
			tmp := l1Ptr.Next
			l1Ptr.Next = l2
			l2 = l2.Next
			l1Ptr.Next.Next = tmp
		}
		l1Ptr = l1Ptr.Next
	}
	//如果l2还有剩余，则补充到l1上
	if l2 != nil {
		l1Ptr.Next = l2
	}

	return l1
}
