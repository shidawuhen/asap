package algorithm

/*
原题：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/
剑指 Offer 43. 1～n 整数中 1 出现的次数
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。



示例 1：

输入：n = 12
输出：5
示例 2：

输入：n = 13
输出：6


限制：

1 <= n < 2^31

分析：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/solution/gui-lu-by-shidawuhen/
*/

func countDigitOne(n int) int {
	div := 10
	sum := 0
	for n/div != 0 {
		residue := n % div
		left := n - residue
		sum += (left / div) * (div / 10)
		if residue/(div/10) >= 2 {
			sum += div / 10
		} else if residue/(div/10) >= 1 {
			sum += (residue + 1 - div/10)
		}
		//fmt.Println(sum)
		div *= 10
	}

	if n/(div/10) >= 2 {
		sum += div / 10
	} else if n/(div/10) >= 1 {
		sum += (n + 1 - div/10)
	}

	return sum
}
