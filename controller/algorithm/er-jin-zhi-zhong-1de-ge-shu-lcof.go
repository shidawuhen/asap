package algorithm

import (
	"bytes"
	"fmt"
)

/*
原题：https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
剑指 Offer 15. 二进制中1的个数
请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。

示例 1：

输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
示例 2：

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
示例 3：

输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。
*/

func HammingWeight(num uint32) int {
	buffer := new(bytes.Buffer)
	fmt.Fprintf(buffer, "%b", num)
	res := buffer.String()
	sum := 0
	for i := 0; i < len(res); i++ {
		if res[i] == '1' {
			sum++
		}
	}
	return sum
}
