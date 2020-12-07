package algorithm

import (
	"fmt"
	"math"
)

/*
原题：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
剑指 Offer 56 - II. 数组中数字出现的次数 II
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。



示例 1：

输入：nums = [3,4,3,3]
输出：4
示例 2：

输入：nums = [9,1,7,9,7,9,7]
输出：1


限制：

1 <= nums.length <= 10000
1 <= nums[i] < 2^31
分析：
每位上的值，正好是3的倍数没问题，不是三的倍数说明是单独的那个值
*/

func singleNumber(nums []int) int {
	count := make([]int, 32)
	for i := 0; i < len(nums); i++ {
		res := fmt.Sprintf("%b", nums[i])
		for j := 0; j < len(res); j++ {
			if res[j] == '1' {
				count[len(res)-1-j]++
			}
		}
	}
	sum := 0
	for i, v := range count {
		count[i] = v % 3
		if count[i] == 1 {
			sum += int(math.Pow(2, float64(count[i])))
		}
	}
	//fmt.Println(count)
	return sum
}
