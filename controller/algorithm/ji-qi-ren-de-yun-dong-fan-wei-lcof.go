package algorithm

import "strconv"

/*
原题：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
剑指 Offer 13. 机器人的运动范围
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外）
，也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？



示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k <= 20
分析：
广度优先遍历
*/

var iToSplit map[int]int

type mapIndex struct {
	x int
	y int
}

func canJoin(x, y, k int) bool {
	if sumInt(x)+sumInt(y) > k {
		return false
	}
	return true
}

func sumInt(x int) int {
	if v, ok := iToSplit[x]; ok {
		return v
	}
	xS := strconv.Itoa(x)
	sum := 0
	for i := 0; i < len(xS); i++ {
		num, _ := strconv.Atoi(string(xS[i]))
		sum += num
	}
	iToSplit[x] = sum
	return sum
}

func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if k == 0 {
		return 1
	}
	iToSplit = make(map[int]int, 0)
	used := make([][]int, m)
	for i := 0; i < m; i++ {
		used[i] = make([]int, n)
	}
	queue := make([]mapIndex, 0)
	queue = append(queue, mapIndex{0, 0})
	used[0][0] = 1
	for i := 0; i < len(queue); i++ {
		mIndex := queue[i].x
		nIndex := queue[i].y
		//上下左右判断能否加入
		//上
		if mIndex-1 >= 0 && used[mIndex-1][nIndex] == 0 && canJoin(mIndex-1, nIndex, k) {
			queue = append(queue, mapIndex{mIndex - 1, nIndex})
			used[mIndex-1][nIndex] = 1
		}
		//下
		if mIndex+1 < m && used[mIndex+1][nIndex] == 0 && canJoin(mIndex+1, nIndex, k) {
			queue = append(queue, mapIndex{mIndex + 1, nIndex})
			used[mIndex+1][nIndex] = 1
		}
		//左
		if nIndex-1 >= 0 && used[mIndex][nIndex-1] == 0 && canJoin(mIndex, nIndex-1, k) {
			queue = append(queue, mapIndex{mIndex, nIndex - 1})
			used[mIndex][nIndex-1] = 1
		}
		//右
		if nIndex+1 < n && used[mIndex][nIndex+1] == 0 && canJoin(mIndex, nIndex+1, k) {
			queue = append(queue, mapIndex{mIndex, nIndex + 1})
			used[mIndex][nIndex+1] = 1
		}
	}
	return len(queue)
}
