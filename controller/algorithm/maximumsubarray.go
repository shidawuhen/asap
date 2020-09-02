package algorithm

/*
原题：
https://leetcode-cn.com/problems/maximum-subarray/
https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/submissions/
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

分析：
解法1：使用分治法，按文章所写方案计算
解法2：从左往右遍历，将遍历的数据求和，如果和小于0，表示这个范围内的都没有用，因为和后面的相加，只会减小和值
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

//分治法
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

//简单解法
//从左往右扫描，如果sum值小于0了，则放弃已经走过的值，重新开始
//[-2,1,-3,4,-1,2,1,-5,4]
func MaxSubArraySimple(nums []int) int {
	sum := 0
	maxSum := 0
	length := len(nums)
	if length == 0 {
		return maxSum
	}
	maxSum = nums[0]
	for i := 0; i < length; i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}
