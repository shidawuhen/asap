package algorithm

/*
原题：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
剑指 Offer 53 - I. 在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0


限制：

0 <= 数组长度 <= 50000

分析：
二分法查找到位置，然后向左向右查找
*/
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	l := 0
	r := length -1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {//说明找到了
			count := 1
			for j := m - 1; j >= l;j--{
				if nums[j] == target{
					count++
				}else{
					break
				}
			}
			for j := m + 1; j <= r; j++ {
				if nums[j] == target{
					count++
				}else{
					break
				}
			}
			return count
		}else if nums[m] < target{
			l = m + 1
		}else{
			r = m - 1
		}
	}
	return 0
}