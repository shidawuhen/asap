package algorithm

/*
原题：https://leetcode-cn.com/problems/intersection-of-two-arrays/
349. 两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]


说明：

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
*/
func intersection(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				record[nums1[i]] = 1
				break
			}
		}
	}
	res := make([]int, len(record))
	i := 0
	for k, _ := range record {
		res[i] = k
		i++
	}
	return res
}
