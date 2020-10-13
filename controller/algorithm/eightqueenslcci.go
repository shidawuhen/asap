package algorithm

import (
	"math"
)

/*
原题：https://leetcode-cn.com/problems/eight-queens-lcci/
面试题 08.12. 八皇后
设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。

注意：本题相对原题做了扩展

示例:

 输入：4
 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 解释: 4 皇后问题存在如下两个不同的解法。
[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]

*/

var res [][]string

func SolveNQueens(n int) [][]string {
	res = make([][]string, 0)
	mark := make([]int, n) //记录放在第几列上
	for i := 0; i < n; i++ {
		mark[0] = i
		computeQueen(mark, n, 1)
	}
	//fmt.Println(res)
	return res
}

func addSoluation(mark []int, n int) {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			if j != mark[i] {
				str += "."
			} else {
				str += "Q"
			}
		}
		s[i] = str
	}
	res = append(res, s)
}

//计算第index皇后，在第index行的位置
func computeQueen(mark []int, n int, index int) {
	if index == n {
		addSoluation(mark, n)
	}
	for i := 0; i < n; i++ {
		if canPlace(mark, index, i) {
			mark[index] = i
			computeQueen(mark, n, index+1)
		}
	}
}

//index放置到place上，和其他
func canPlace(mark []int, index int, place int) bool {
	for i := 0; i < index; i++ {
		if mark[i] == place || //同一列或者同一对角线(两个边长相等)
			int(math.Abs(float64(place)-float64(mark[i]))) == int(math.Abs(float64(index)-float64(i))) {
			return false
		}
	}
	return true
}
