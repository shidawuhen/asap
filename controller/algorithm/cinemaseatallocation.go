package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/cinema-seat-allocation/
1386. 安排电影院座位


如上图所示，电影院的观影厅中有 n 行座位，行编号从 1 到 n ，且每一行内总共有 10 个座位，列编号从 1 到 10 。

给你数组 reservedSeats ，包含所有已经被预约了的座位。比如说，researvedSeats[i]=[3,8] ，它表示第 3 行第 8 个座位被预约了。

请你返回 最多能安排多少个 4 人家庭 。4 人家庭要占据 同一行内连续 的 4 个座位。隔着过道的座位（比方说 [3,3] 和 [3,4]）不是连续的座位，但是如果你可以将 4 人家庭拆成过道两边各坐 2 人，这样子是允许的。



示例 1：



输入：n = 3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
输出：4
解释：上图所示是最优的安排方案，总共可以安排 4 个家庭。蓝色的叉表示被预约的座位，橙色的连续座位表示一个 4 人家庭。
示例 2：

输入：n = 2, reservedSeats = [[2,1],[1,8],[2,6]]
输出：2
示例 3：

输入：n = 4, reservedSeats = [[4,3],[1,4],[4,6],[1,7]]
输出：4


提示：

1 <= n <= 10^9
1 <= reservedSeats.length <= min(10*n, 10^4)
reservedSeats[i].length == 2
1 <= reservedSeats[i][0] <= n
1 <= reservedSeats[i][1] <= 10
所有 reservedSeats[i] 都是互不相同的。

分析:
每行最多能放两个4人家庭，每行最多有三个位置适合放。三个位置为2345，4567，6789
从左往右放置，看看最多能放置多少

*/
func canSeat(temp []int, a, b, c, d int) bool {
	if temp[a] == 0 && temp[b] == 0 && temp[c] == 0 && temp[d] == 0 {
		temp[a], temp[b], temp[c], temp[d] = 1, 1, 1, 1
		return true
	}
	return false
}

func MaxNumberOfFamilies(n int, reservedSeats [][]int) int {
	sum := 0
	if n == 0 {
		return sum
	}
	seatMap := make(map[int][]int)
	rl := len(reservedSeats)
	for i := 0; i < rl; i++ {
		seatMap[reservedSeats[i][0]] = append(seatMap[reservedSeats[i][0]], reservedSeats[i][1])
	}
	mapLen := len(seatMap)
	sum = (n - mapLen) * 2
	for _, arr := range seatMap {
		temp := make([]int, 11)
		for j := 0; j < len(arr); j++ {
			temp[arr[j]] = 1
		}
		p1, p2, p3 := false, false, false
		p1 = canSeat(temp, 2, 3, 4, 5)
		if p1 {
			sum++
		}
		if !p1 {
			p2 = canSeat(temp, 4, 5, 6, 7)
			if p2 {
				sum++
			}
		}
		if !p2 {
			p3 = canSeat(temp, 6, 7, 8, 9)
			if p3 {
				sum++
			}
		}
	}
	fmt.Println(sum)
	return sum
}
