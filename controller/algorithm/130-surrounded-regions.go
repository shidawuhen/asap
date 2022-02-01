/**
@author: Jason Pang
@desc:
@date: 2022/2/1
**/
package algorithm

type BoardPoint struct {
	row, col int
}

func solve(board [][]byte) {
	row := len(board)
	col := len(board[0])
	used := make([][]bool, row)
	for i := 0; i < row; i++ {
		used[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' && used[i][j] == false { //是O且没有用过
				oList := make([]BoardPoint, 0)
				oList = append(oList, BoardPoint{
					i,
					j,
				})
				used[i][j] = true
				change := true
				if oList[0].row == 0 || oList[0].row == row-1 || oList[0].col == 0 || oList[0].col == col-1 {
					change = false
				}
				//遍历oList，找出临近未用过的点，判断是否可以改变
				for k := 0; k < len(oList); k++ {
					r := oList[k].row
					c := oList[k].col
					//上
					if r-1 >= 0 && board[r-1][c] == 'O' && used[r-1][c] == false {
						oList = append(oList, BoardPoint{
							r - 1,
							c,
						})
						used[r-1][c] = true
						if r-1 == 0 {
							change = false
						}
					}
					//下
					if r+1 < row && board[r+1][c] == 'O' && used[r+1][c] == false {
						oList = append(oList, BoardPoint{
							r + 1,
							c,
						})
						used[r+1][c] = true
						if r+1 == row-1 {
							change = false
						}
					}
					//左
					if c-1 >= 0 && board[r][c-1] == 'O' && used[r][c-1] == false {
						oList = append(oList, BoardPoint{
							r,
							c - 1,
						})
						used[r][c-1] = true
						if c-1 == 0 {
							change = false
						}
					}
					//右
					if c+1 < col && board[r][c+1] == 'O' && used[r][c+1] == false {
						oList = append(oList, BoardPoint{
							r,
							c + 1,
						})
						used[r][c+1] = true
						if c+1 == col-1 {
							change = false
						}
					}
				}
				if change {
					for _, item := range oList {
						board[item.row][item.col] = 'X'
					}
				}
			}
		}
	}
}
