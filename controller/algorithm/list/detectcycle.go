/**
@author: Jason Pang
@desc:
@date: 2022/8/4
**/
package list

/**
这种题，如果提前没有思路的话，应该很难想出来。使用快慢指针，需要做证明，为什么快指针和慢指针肯定能碰上。
有了思路后，再做就容易多了
方向就是，你就不断的走就行，要是快指针到nil了，那就是没有，否则肯定会碰上，不要怕，就是干。

用乘法，不用除法；用加法不用减法
能自己判断出会相遇，但是相遇之后怎么做没搞出来，还是因为上面思旭错了
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	ptr := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}
	return nil
}
