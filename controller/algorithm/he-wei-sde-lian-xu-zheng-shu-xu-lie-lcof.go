package algorithm

import (
	"math"
)

/*
原题：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
剑指 Offer 57 - II. 和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]


限制：

1 <= target <= 10^5

分析：
公式=n(n+1)/2-a(a-1)/2
n为最後的数，a为最前的数
*/

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for a := 1; a <= target/2; a++ {
		n := (-1 + math.Sqrt(float64(1+8*target+4*a*(a-1)))) / 2
		if math.Ceil(n)-math.Floor(n) < 0.000001 {
			d := make([]int, 0)
			for i := a; i <= int(n); i++ {
				d = append(d, i)
			}
			res = append(res, d)
		}
	}
	return res
}
