package algorithm

import "sort"

/*
原题：https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
剑指 Offer 61. 扑克牌中的顺子
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。



示例 1:

输入: [1,2,3,4,5]
输出: True


示例 2:

输入: [0,0,1,2,5]
输出: True


限制：

数组长度为 5

数组的数取值为 [0, 13] .

分析：
先排序，然后把0的去掉，剩下的看差值是不是小于0的个数

*/

func isStraight(nums []int) bool {
	if len(nums) < 5 {
		return false
	}
	sort.Ints(nums)
	zeroNum := 0
	durSum := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zeroNum++
		} else {
			if nums[i+1]-nums[i] == 0 {
				return false
			}
			durSum += nums[i+1] - nums[i] - 1
		}
	}
	if durSum > zeroNum {
		return false
	}
	return true
}
