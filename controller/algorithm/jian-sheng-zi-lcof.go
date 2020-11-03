package algorithm

/*
原题：https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
剑指 Offer 14- I. 剪绳子
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
提示：

2 <= n <= 58

分析：
使用动态规划，第一次知道，动态规划可以不用来计算所有数值
*/

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]int64, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 0; j <= i/2; j++ {
			dp[i] = Max(dp[i], dp[j]*dp[i-j])
		}
	}
	return int(dp[n])
}
