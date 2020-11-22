package algorithm

/*
原题：https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
剑指 Offer 39. 数组中出现次数超过一半的数字
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。



你可以假设数组是非空的，并且给定的数组总是存在多数元素。



示例 1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2


限制：

1 <= 数组长度 <= 50000

分析：
1.用快排，然后找到中间位置
2.用hashmap
*/

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//1.快排
	/*quickSort(nums,0,len(nums)-1)
	return nums[len(nums)/2]*/
	//2.hashmap
	record := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		record[nums[i]]++
		if record[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
