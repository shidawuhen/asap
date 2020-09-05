package algorithm

/*
原题：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
搜索二维矩阵
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。

分析：
每次选择最中心的点，如果一致，则找到，如果不一致，判断四分后的每块位置是否可能有该值，在范围之内的则继续遍历，直到矩阵为1*1结束
注意点:
1.middle的值必须变动，否则容易死循环
2.算出middle后，对middle的加减需要判断边界
*/

//(ly,lx) (ry,rx)为坐标
func findTarget(matrix [][]int, target int, ly int, lx int, ry int, rx int) bool {
	if lx > rx || ly > ry {
		return false
	}

	mx := (lx + rx) / 2
	my := (ly + ry) / 2
	//如果中间位置的值和target一致，则返回
	mv := matrix[my][mx]
	if mv == target {
		return true
	}
	//判断四个区域的范围是否有可能存在该值
	//左上符合条件
	if my-1 >= ly && matrix[ly][lx] <= target && matrix[my-1][mx] >= target {
		res := findTarget(matrix, target, ly, lx, my-1, mx)
		if res == true {
			return true
		}
	}
	//右上符合条件
	if mx+1 <= rx && matrix[ly][mx+1] <= target && matrix[my][rx] >= target {
		res := findTarget(matrix, target, ly, mx+1, my, rx)
		if res == true {
			return true
		}
	}
	//左下角符合条件
	if mx-1 >= lx && matrix[my][lx] <= target && matrix[ry][mx-1] >= target {
		res := findTarget(matrix, target, my, lx, ry, mx-1)
		if res == true {
			return true
		}
	}
	//右下角符合条件
	if my+1 <= ry && matrix[my+1][mx] <= target && matrix[ry][rx] >= target {
		res := findTarget(matrix, target, my+1, mx, ry, rx)
		if res == true {
			return true
		}
	}
	return false
}

func SearchMatrix(matrix [][]int, target int) bool {
	res := false
	if len(matrix) == 0 {
		return res
	}
	xl := len(matrix[0])
	yl := len(matrix)

	res = findTarget(matrix, target, 0, 0, yl-1, xl-1)
	return res
}
