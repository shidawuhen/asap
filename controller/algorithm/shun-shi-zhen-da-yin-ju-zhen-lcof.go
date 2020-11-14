package algorithm

/*
原题：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
剑指 Offer 29. 顺时针打印矩阵
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100

分析：
这道题主要困难在于边界判断
*/

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	res := make([]int, row*col)
	tIndex := 0
	bIndex := row - 1
	lIndex := 0
	rIndex := col - 1
	index := 0
	for !(tIndex > bIndex || lIndex > rIndex) {
		//fmt.Println(tIndex,bIndex,lIndex,rIndex,res)
		//顶行
		//if lIndex < col && rIndex >= 0 && tIndex < row {
		for i := lIndex; i <= rIndex; i++ {
			res[index] = matrix[tIndex][i]
			index++
		}
		//}
		tIndex++
		if tIndex > bIndex {
			return res
		}
		//右列
		//if tIndex < row && bIndex >= 0 && rIndex >= 0 {
		for i := tIndex; i <= bIndex; i++ {
			res[index] = matrix[i][rIndex]
			index++
		}
		//}
		rIndex--
		if rIndex < lIndex {
			return res
		}
		//下行
		//if rIndex >= 0 && lIndex < col && bIndex >= 0 {
		for i := rIndex; i >= lIndex; i-- {
			res[index] = matrix[bIndex][i]
			index++
		}
		//}
		bIndex--
		//fmt.Println(bIndex,tIndex)
		if bIndex < tIndex {
			return res
		}
		//左列
		//if bIndex >= 0 && tIndex < row && lIndex < col {
		for i := bIndex; i >= tIndex; i-- {
			res[index] = matrix[i][lIndex]
			index++
		}
		//}
		lIndex++
		if lIndex > rIndex {
			return res
		}
	}
	return res
}
