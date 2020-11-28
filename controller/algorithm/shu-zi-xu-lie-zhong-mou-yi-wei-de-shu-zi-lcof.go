package algorithm

import (
	"math"
	"strconv"
)

/*
原题：https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
剑指 Offer 44. 数字序列中某一位的数字
数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。



示例 1：

输入：n = 3
输出：3
示例 2：

输入：n = 11
输出：0


限制：

0 <= n < 2^31

分析：
https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/gui-lu-by-shidawuhen-2/
*/
func findNthDigit(n int) int {
	if n < 10 {
		return n
	} else {
		n = n - 10
		num := 90
		l := 2
		for n-num*l >= 0 {
			n -= num * l
			num *= 10
			l++
		}
		div := n / l
		reminder := n % l
		start := int(math.Pow(10, float64(l-1)))
		start += div
		//找出start的第reminder的值
		startString := strconv.Itoa(start)
		v, _ := strconv.Atoi(string(startString[reminder]))
		return v
	}
}
