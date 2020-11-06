package algorithm

/*
原题：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
剑指 Offer 16. 数值的整数次方
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。



示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25


说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

分析：
蛮力法会超时，使用分治法，同时做记录。有点动态规划的味道在

*/

var powRecord map[int]float64

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	powRecord = make(map[int]float64)
	y := n
	if n < 0 {
		y = (-1) * n
	}
	res := computePow(x, y)
	if n < 0 {
		return 1 / res
	} else {
		return res
	}
}

func computePow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	l, r := 0.0, 0.0
	if v, ok := powRecord[n/2]; ok {
		l = v
	} else {
		l = computePow(x, n/2)
		powRecord[n/2] = l
	}

	if v, ok := powRecord[n-n/2]; ok {
		r = v
	} else {
		r = computePow(x, n-n/2)
		powRecord[n-n/2] = r
	}
	return l * r
}
