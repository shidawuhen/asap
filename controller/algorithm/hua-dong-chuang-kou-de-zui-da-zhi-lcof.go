package algorithm

import "container/list"

/*
原题：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
剑指 Offer 59 - I. 滑动窗口的最大值
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

分析：
使用队列，构建非严格递减队列
这个思路还是很厉害的
*/

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 {
		return nums
	}
	res := make([]int, len(nums)-k+1)
	queue := list.New()
	//queue.PushBack(nums[0])
	for i := 0; i < len(nums); i++ {
		e := queue.Front()
		if i-k >= 0 && nums[i-k] == e.Value.(int) {
			queue.Remove(e)
			e = queue.Front()
		}
		e = queue.Back()
		for e != nil && e.Value.(int) < nums[i] {
			queue.Remove(e)
			e = queue.Back()
		}
		queue.PushBack(nums[i])
		if i-k+1 >= 0 {
			res[i-k+1] = queue.Front().Value.(int)
		}
	}
	return res
}
