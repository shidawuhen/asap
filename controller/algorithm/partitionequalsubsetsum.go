package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/partition-equal-subset-sum/
416. 分割等和子集
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].


示例 2:

输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.

分析：
解法一：
使用回溯法，并且记录同一层的remain，用于剪枝
*/

func CanPartition(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	record := make(map[int]map[int]int)
	remain := sum / 2
	res := doPartition(0, remain, nums, record)
	fmt.Println(res)
	return res
}

func doPartition(level int, remain int, nums []int, record map[int]map[int]int) bool {
	//fmt.Println(level,remain)
	if remain == 0 {
		return true
	}
	if remain < 0 || level >= len(nums) {
		return false
	}
	if _, ok := record[level][remain]; ok {
		return false
	}
	if _, ok := record[level]; !ok {
		record[level] = make(map[int]int)
	}
	record[level][remain] = 1
	return doPartition(level+1, remain-nums[level], nums, record) || doPartition(level+1, remain, nums, record)

}
