package algorithm

/*
原题：https://leetcode-cn.com/problems/chou-shu-lcof/
剑指 Offer 49. 丑数
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。



示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。
分析：动态规划
这道题解法不是自己想出来的。没有从动态规划的思路去考虑，其实挺不应该的
*/
func nthUglyNumber(n int) int {
	record := make([]int, n)
	a, b, c := 0, 0, 0
	record[0] = 1
	for i := 1; i < n; i++ {
		av := record[a] * 2
		bv := record[b] * 3
		cv := record[c] * 5
		if av <= bv && av <= cv {
			record[i] = av
			a++
		}
		if bv <= av && bv <= cv {
			record[i] = bv
			b++
		}
		if cv <= av && cv <= bv {
			record[i] = cv
			c++
		}
	}
	//fmt.Println(record)
	return record[n-1]
}
