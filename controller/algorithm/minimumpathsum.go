package algorithm

import (
	"container/list"
)

/*
原题：https://leetcode-cn.com/problems/minimum-path-sum/
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
分析：
解法一：本题使用标准动态规划算法，超时
1. 如果一条路径最短，那么其子集必定最短，否则矛盾
2. 公式：d(ij,Jpq)=min{Cij-pq,d(pq)) 表示如果从ij经过以pq作为左上角的矩阵，到达终点的最短距离
3. 记录数据：用二维数组记，ij表示从ij顶点到终点的最短距离。由底向上计算
解法二：本题使用灵活的动态规划算法
使用解法一，充分说明了什么叫按照固定模式走是打脸。我还用了队列，时间复杂度也是O(m*n),不过是2*m*n，空间也多用了一倍
这道题其实有更简单的方法，不适用自底向上，使用自上到下。思路为，从顶点到某个点的最短距离，就是从顶点到该点（上方/左方）最短距离加上该点的距离。
连空间都能复用
*/

type point struct {
	x int
	y int
}
type s struct {
	mainPoint point
	matrix    point
}

//标准方案：超时
func MinPathSum(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	//申请空间
	valMap := make([][]int, row)
	for i := 0; i < row; i++ {
		valMap[i] = make([]int, col)
	}
	//使用先进先出队列，从终点反推到对应点的最短距离
	queue := list.New()
	queue.PushBack(s{mainPoint: point{row - 1, col - 1}, matrix: point{row - 1, col - 1}})
	e := queue.Front()
	for e != nil {
		//获取头部
		v := e.Value.(s)
		//计算该点通过对应矩阵到达终点的最短距离
		cx := v.mainPoint.x
		cy := v.mainPoint.y
		minDis := grid[cx][cy]
		if valMap[v.matrix.x][v.matrix.y] != -1 {
			minDis += valMap[v.matrix.x][v.matrix.y]
		}
		if valMap[cx][cy] == 0 || minDis < valMap[cx][cy] {
			valMap[cx][cy] = minDis
		}
		//计算完毕后，将上方和左侧节点放入
		if cx-1 >= 0 { //上
			queue.PushBack(s{mainPoint: point{cx - 1, cy}, matrix: point{cx, cy}})
		}
		if cy-1 >= 0 {
			queue.PushBack(s{mainPoint: point{cx, cy - 1}, matrix: point{cx, cy}})
		}
		queue.Remove(e)
		e = queue.Front()
	}
	//fmt.Println(valMap,valMap[0][0])
	return valMap[0][0]
}

func MinPathSumSimple(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			minDis := 0
			if i-1 >= 0 && (grid[i-1][j] < minDis || minDis == 0) { //上
				minDis = grid[i-1][j]
			}

			if j-1 >= 0 && (grid[i][j-1] < minDis || minDis == 0) { //左
				minDis = grid[i][j-1]
			}
			grid[i][j] = minDis + grid[i][j]
		}
	}
	//fmt.Println(grid,grid[row-1][col-1])
	return grid[row-1][col-1]
}

//例题2 剑指 Offer 47. 礼物的最大价值 https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
func maxValue(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			minDis := 0
			if i-1 >= 0 && (grid[i-1][j] > minDis || minDis == 0) { //上
				minDis = grid[i-1][j]
			}

			if j-1 >= 0 && (grid[i][j-1] > minDis || minDis == 0) { //左
				minDis = grid[i][j-1]
			}
			grid[i][j] = minDis + grid[i][j]
		}
	}
	//fmt.Println(grid,grid[row-1][col-1])
	return grid[row-1][col-1]
}
