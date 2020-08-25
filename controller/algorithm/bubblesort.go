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

分析：练习排序算法，此处使用冒泡法
使用改进后的冒泡排序，但是效率还是太低，超时了
*/

func Exchange(nums []int) []int {
	length := len(nums)
	exchange := length - 1
	oddeven := make([]int, length)
	for i := 0; i < length; i++ {
		if nums[i]%2 == 0 {
			oddeven[i] = 0
		} else {
			oddeven[i] = 1
		}
	}
	for exchange != 0 {
		boundary := exchange
		exchange = 0
		for j := 0; j < boundary; j++ {
			if oddeven[j] == 0 && oddeven[j+1] != 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				oddeven[j], oddeven[j+1] = oddeven[j+1], oddeven[j]
				exchange = j
			}
		}
	}
	return nums
}
