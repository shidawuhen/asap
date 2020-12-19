package algorithm

/*
原题：https://leetcode-cn.com/problems/qiu-12n-lcof/
剑指 Offer 64. 求1+2+…+n
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。



示例 1：

输入: n = 3
输出: 6
示例 2：

输入: n = 9
输出: 45


限制：

1 <= n <= 10000
分析：用递归代替循环
*/
var allSumNums int

func sumNums(n int) int {
	allSumNums = 0
	addNums(n)
	return allSumNums
}

func addNums(n int) bool {
	allSumNums += n
	return n > 0 && addNums(n-1)
}
