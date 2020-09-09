package algorithm

/*
原题：https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table/
668. 乘法表中第k小的数
几乎每一个人都用 乘法表。但是你能在乘法表中快速找到第k小的数字吗？

给定高度m 、宽度n 的一张 m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。

例 1：

输入: m = 3, n = 3, k = 5
输出: 3
解释:
乘法表:
1	2	3
2	4	6
3	6	9

第5小的数字是 3 (1, 2, 2, 3, 3).
例 2：

输入: m = 2, n = 3, k = 6
输出: 6
解释:
乘法表:
1	2	3
2	4	6

第6小的数字是 6 (1, 2, 2, 3, 4, 6).
注意：

m 和 n 的范围在 [1, 30000] 之间。
k 的范围在 [1, m * n] 之间。

分析：
使用折半法。计算最小最大的中间值，然后判断小于该值的数量，如果大于k，则用最小值和中间值再求中间值，如果小于k，则用中间值和最大值再求中间值
需要考虑如何计算小于该值的数量
最大最小的中间值未必是该数组中有的值，如何确保最终值为数组中的值呢？
1. 通过start和end最终值变为一致
2. start是不断加一的
3. count >= k的时候，end=middle
4. 记住一个前提，只有数组中的值才会在k上
*/

func FindKthNumber(m int, n int, k int) int {
	if n == 0 || m == 0 {
		return 0
	}
	count := 0
	start := 1
	end := m*n
	middleV := 0
	for start < end {
		//从左下角开始，如果太大，则m减小，如果太小，则n增大
		middleV = (start + end)/2
		//计算小于等于middleV的数量
		count = 0
		row := m
		col := 1
		for row >=1 && col <= n {
			if row*col <= middleV{
				count += row
				col++
			}else{
				row--
			}
		}
		//需要往大的移动
		if count < k {
			start = middleV + 1
		}
		//需要往小的移动
		if count >= k {
			end = middleV
		}
	}
	return start
}
