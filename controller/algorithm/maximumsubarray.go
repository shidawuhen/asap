package algorithm

/*
原题：https://leetcode-cn.com/problems/maximum-subarray/
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

分析：
解法1：使用分治法，按文章所写方案计算
*/

func maxSum(nums []int, startIndex int, endIndex int) int {
	if startIndex >= endIndex {
		return nums[endIndex]
	}
	middleIndex := (startIndex + endIndex) / 2
	//计算左侧的最大子段和
	lSum := maxSum(nums, startIndex, middleIndex)
	//计算右侧的最大子段和
	rSum := maxSum(nums, middleIndex+1, endIndex)
	//计算中间的最大子段和。求middleIndex左侧和右侧的最大值
	leftSum := 0
	leftMaxSum := nums[middleIndex]
	for i := middleIndex; i >= 0; i-- {
		leftSum += nums[i]
		if leftSum > leftMaxSum {
			leftMaxSum = leftSum
		}
	}
	rightSum := 0
	rightMaxSum := nums[endIndex]
	if middleIndex+1 < endIndex {
		rightMaxSum = nums[middleIndex+1]
	}
	for i := middleIndex + 1; i <= endIndex; i++ {
		rightSum += nums[i]
		if rightSum > rightMaxSum {
			rightMaxSum = rightSum
		}
	}
	sum := leftMaxSum + rightMaxSum
	if lSum > sum {
		sum = lSum
	}
	if rSum > sum {
		sum = rSum
	}
	return sum
}

//[-2,1,-3,4,-1,2,1,-5,4]
func MaxSubArray(nums []int) int {
	sum := 0
	length := len(nums)
	if length == 0 {
		return sum
	}
	sum = maxSum(nums, 0, length-1)
	return sum
}
