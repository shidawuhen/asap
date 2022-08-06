/**
@author: Jason Pang
@desc:
@date: 2022/8/5
**/
package list

/**
删除倒数节点，这种就是干
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	ahead := head
	for ahead.Next != nil && n > 0 {
		ahead = ahead.Next
		n--
	}
	if n > 0 {
		head = head.Next
		return head
	}
	bhead := head
	//head和ahead一起移动
	for ahead.Next != nil {
		ahead = ahead.Next
		bhead = bhead.Next
	}
	//删除
	if bhead.Next == nil {
		return head
	}
	bhead.Next = bhead.Next.Next
	return head
}
