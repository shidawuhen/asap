package algorithm

/*
原题：https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
剑指 Offer 62. 圆圈中最后剩下的数字
0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。

例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。



示例 1：

输入: n = 5, m = 3
输出: 3
示例 2：

输入: n = 10, m = 17
输出: 2


限制：

1 <= n <= 10^5
1 <= m <= 10^6

分析：
长度为n的数据，将第一个数去掉之后，其实变为了长度为n-1的数组，n-1的数组最后剩的位置为x，则长度为n的数组剩的位置为m+x
*/
func lastRemaining(n int, m int) int {
	return lastRemainingF(n, m)
}

func lastRemainingF(n int, m int) int {
	if n == 1 {
		return 0
	}
	x := lastRemainingF(n-1, m)
	return (m + x) % n
}
