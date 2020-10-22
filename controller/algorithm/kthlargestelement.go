package algorithm

/*
原题：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
215. 数组中的第K个最大元素
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

分析：
借鉴快排的划分过程，计算出轴值后，判断轴值与k的大小，选择合适区间

*/

func FindKthLargest(nums []int, k int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	k--
	startIndex := 0
	endIndex := length - 1
	index := quicksort(nums, startIndex, endIndex)
	for index != k {
		if index < k {
			startIndex = index + 1
		} else {
			endIndex = index - 1
		}
		index = quicksort(nums, startIndex, endIndex)
	}
	//fmt.Println(k,nums[index])
	return nums[index]
}

func quicksort(nums []int, startIndex int, endIndex int) int {
	for startIndex < endIndex {
		for startIndex < endIndex {
			if nums[endIndex] <= nums[startIndex] {
				endIndex--
			} else {
				nums[startIndex], nums[endIndex] = nums[endIndex], nums[startIndex]
				startIndex++
				break
			}
		}
		for startIndex < endIndex {
			if nums[startIndex] >= nums[endIndex] {
				startIndex++
			} else {
				nums[startIndex], nums[endIndex] = nums[endIndex], nums[startIndex]
				endIndex--
				break
			}
		}
	}
	return startIndex
}
