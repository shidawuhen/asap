package algorithm

/*
原题：https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
剑指 Offer 65. 不用加减乘除做加法
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。



示例:

输入: a = 1, b = 1
输出: 2


提示：

a, b 均可能是负数或 0
结果不会溢出 32 位整数

分析：
位运算
将加法分为两部分：进位之后的值，不进位的值。
b一直记录进位的值，a记录不进位的值，当b没有进位的时候，表示已经结束了

*/

func add(a int, b int) int {
	// 进位
	var carry int
	for b != 0 {
		// 进位
		carry = (a & b) << 1
		// 不进位加
		a ^= b
		// 加进位
		b = carry
	}
	return a
}
