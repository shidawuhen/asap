package algorithm

/*
原题：https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。


提示：
1 <= nums.length <= 50000
1 <= nums[i] <= 10000

分析：练习二路归并算法
二路归并算法在合并的时候，要注意回写。这个算法时间复杂度在O(nlog2^n)
*/

func Merge(nums []int, res []int, startIndex int, middleIndex int, endIndex int) {
	i := startIndex
	j := middleIndex + 1
	if j > endIndex {
		j = endIndex
	}
	k := startIndex

	change := 1
	for change == 1 {
		change = 0
		if i <= middleIndex && nums[i]%2 == 1 {
			res[k] = nums[i]
			k++
			i++
			change = 1
		}
		if j <= endIndex && nums[j]%2 == 1 {
			res[k] = nums[j]
			k++
			j++
			change = 1
		}
	}

	for i <= middleIndex {
		res[k] = nums[i]
		k++
		i++
	}
	for j <= endIndex {
		res[k] = nums[j]
		k++
		j++
	}
	//回写
	for i := startIndex; i <= endIndex; i++ {
		nums[i] = res[i]
	}
}

func MergeSort(nums []int, res []int, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		res[endIndex] = nums[endIndex]
		return
	}
	middleIndex := (startIndex + endIndex) / 2
	MergeSort(nums, res, startIndex, middleIndex)
	MergeSort(nums, res, middleIndex+1, endIndex)
	Merge(nums, res, startIndex, middleIndex, endIndex)
	return
}

//二路归并
func ExchangeUseMerge(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nums
	}
	res := make([]int, length)
	MergeSort(nums, res, 0, length-1)
	return res
}

//算法二：简单版
func ExchangeSimple(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nums
	}
	res := make([]int, length)
	i := 0
	j := length - 1
	for k := 0; k < length; k++ {
		if nums[k]%2 == 1 {
			res[i] = nums[k]
			i++
		} else {
			res[j] = nums[k]
			j--
		}
	}
	return res
}
