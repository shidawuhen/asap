package algorithm

/*
原题：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
剑指 Offer 12. 矩阵中的路径
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。



示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
提示：

1 <= board.length <= 200
1 <= board[i].length <= 200

分析：
使用深度优先遍历
*/

func Exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0 {
		return false
	}
	col := len(board[0])
	used := make([][]int, row)
	for i := 0; i < row; i++ {
		used[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res := wordExists(board, i, j, word, 0, used)
			if res == true {
				return true
			}
		}
	}
	return false
}

func wordExists(board [][]byte, row int, col int, word string, index int, used [][]int) bool {
	if len(word) == index {
		return true
	}
	if row >= 0 && row < len(board) && col >= 0 && col < len(board[0]) && board[row][col] == word[index] && used[row][col] == 0 {
		//四个方向是否有合适的
		used[row][col] = 1
		r1, r2, r3, r4 := false, false, false, false
		//上
		//if row - 1 >= 0 {
		r1 = wordExists(board, row-1, col, word, index+1, used)
		if r1 == true {
			return true
		}
		//}
		//下
		//if row + 1 < len(board){
		r2 = wordExists(board, row+1, col, word, index+1, used)
		if r2 == true {
			return true
		}
		//}
		//左
		//if col - 1 >= 0 {
		r3 = wordExists(board, row, col-1, word, index+1, used)
		if r3 == true {
			return true
		}
		//}
		//右
		//if col + 1 < len(board[0]){
		r4 = wordExists(board, row, col+1, word, index+1, used)
		if r4 == true {
			return true
		}
		//}
		used[row][col] = 0
	}
	return false
}
