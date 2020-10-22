package algorithm

/*
原题：https://leetcode-cn.com/problems/triangle/
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。



例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：
如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

分析：
1. 满足最优性原理
2. 公式d(i) = min{vi + d(j)}，d(i)为从i到终点最短距离，j为可以和i相连的点。将二维数组一维化
3. 从底向上计算

*/

//计算该点在一维中的位置，row表示第几行，j表示该行的第几个位置，都从0起
func computePos(row int, j int) int {
	return row*(row+1)/2 + j
}

func MinimumTotal(triangle [][]int) int {
	row := len(triangle)
	if row == 0 {
		return 0
	}
	//从后往上计算
	totalLen := row * (row + 1) / 2
	d := make([]int, totalLen)
	for i := row - 1; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			//判断点j到终点的最短距离
			//1.当前j所处的一维位置
			minDisL1 := 0
			minDisL2 := 0
			pos := computePos(i, j)
			//判断下层相邻的两个点
			l1 := computePos(i+1, j)
			if l1 < totalLen {
				minDisL1 = d[l1]
			}
			l2 := computePos(i+1, j+1)
			if l2 < totalLen {
				minDisL2 = d[l2]
			}
			minDis := minDisL1
			if minDis > minDisL2 {
				minDis = minDisL2
			}
			d[pos] = minDis + triangle[i][j]
		}
	}
	//fmt.Println(d)
	return d[0]
}
