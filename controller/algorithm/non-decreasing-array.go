package algorithm

/*
原题：https://leetcode-cn.com/problems/non-decreasing-array/
665. 非递减数列
给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。



示例 1:

输入: nums = [4,2,3]
输出: true
解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
示例 2:

输入: nums = [4,2,1]
输出: false
解释: 你不能在只改变一个元素的情况下将其变为非递减数列。


说明：

1 <= n <= 10 ^ 4
- 10 ^ 5 <= nums[i] <= 10 ^ 5

特殊用例：
[-1 4 2 3]
[3 4 2 3]
[1 4 1 2]

分析：
贪心法：将导致问题的值改为最小的
*/

func checkPossibility(nums []int) bool {
	//判断有几个不符合排序的.遍历碰到不合适的，就改变
	if len(nums) == 0 {
		return true
	}
	sum := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] { //不匹配的位置
			sum++
			if i == 0 {
				if nums[0] < nums[1] {
					nums[1] = nums[0]
				} else {
					nums[0] = nums[1]
				}
			} else { //如果i-1的值小于等于i+1的值，将i改为i-1的值。否则将i+1改为i的值。贪心法：将值改为最小的
				if nums[i-1] <= nums[i+1] {
					nums[i] = nums[i-1]
				} else {
					nums[i+1] = nums[i]
				}
			}
		}
		if sum > 1 {
			return false
		}
	}
	return true
}
