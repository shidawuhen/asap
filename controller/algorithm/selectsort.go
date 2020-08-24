package algorithm

/*
原题：https://leetcode-cn.com/problems/sort-an-array/
给你一个整数数组 nums，请你将该数组升序排列。

示例 1：

输入：nums = [5,2,3,1]
输出：[1,2,3,5]
示例 2：

输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]
 

提示：

1 <= nums.length <= 50000
-50000 <= nums[i] <= 50000

分析：练习排序算法，此处使用选择排序
结果：测试用例通过，不过超时，O(n^2)还是太高了一些
*/

func SortArray(nums []int) []int {
	length := len(nums)
	for i := 0; i < length - 1; i++ {
		minNum := nums[i]
		index := i
		for j := i + 1;  j < length; j++ {
			if minNum > nums[j] {
				minNum = nums[j]
				index = j
			}
		}
		if index != i {
			nums[i], nums[index] = nums[index],nums[i]
		}
	}
	return nums
}