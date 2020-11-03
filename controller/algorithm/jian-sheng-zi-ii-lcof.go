package algorithm
/*
原题：https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/
剑指 Offer 14- II. 剪绳子 II
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。



示例 1：

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36


提示：

2 <= n <= 1000
*/
import (
	"math/big"
)

func CuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}

	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(0)
	dp[1] = big.NewInt(1)
	dp[2] = big.NewInt(2)
	dp[3] = big.NewInt(3)
	for i := 4; i <= n; i++ {
		dp[i] = big.NewInt(0)
		for j := 1; j <= i/2; j++ {
			d := big.NewInt(1)
			//fmt.Println(i,dp[i],dp[j],dp[i-j],11)
			dp[i] = MaxBig(dp[i], d.Mul(dp[j], dp[i-j]))
		}
	}
	//fmt.Println(dp)
	d := dp[n].Mod(dp[n], big.NewInt(1000000007))
	return int(d.Int64())
}
func MaxBig(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}
