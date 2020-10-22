package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
4. 寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000


提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	index1, index2 := 0, 0
	if (l1+l2)%2 == 1 {
		index1, index2 = (l1+l2)/2, (l1+l2)/2
	} else {
		index1, index2 = (l1+l2)/2-1, (l1+l2)/2
	}
	var res1, res2 int = 0, 0

	i, j := 0, 0
	for i < l1 && j < l2 && (i+j <= index2) {
		if nums1[i] < nums2[j] {
			if i+j == index1 {
				res1 = nums1[i]
			}
			if i+j == index2 {
				res2 = nums1[i]
			}
			i++
		} else {
			if i+j == index1 {
				res1 = nums2[j]
			}
			if i+j == index2 {
				res2 = nums2[j]
			}
			j++
		}
	}
	for i < l1 && (i+j <= index2) {
		if i+j == index1 {
			res1 = nums1[i]
		}
		if i+j == index2 {
			res2 = nums1[i]
		}
		i++
	}
	for j < l2 && (i+j <= index2) {
		if i+j == index1 {
			res1 = nums2[j]
		}
		if i+j == index2 {
			res2 = nums2[j]
		}
		j++
	}
	fmt.Println(index1, index2, i, j, res1, res2, (float64(res1)+float64(res2))/2)
	return (float64(res1) + float64(res2)) / 2
}
