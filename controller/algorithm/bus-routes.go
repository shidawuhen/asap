package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/bus-routes/
815. 公交路线
我们有一系列公交路线。每一条路线 routes[i] 上都有一辆公交车在上面循环行驶。例如，有一条路线 routes[0] = [1, 5, 7]，
表示第一辆 (下标为0) 公交车会一直按照 1->5->7->1->5->7->1->... 的车站路线行驶。

假设我们从 S 车站开始（初始时不在公交车上），要去往 T 站。
期间仅可乘坐公交车，求出最少乘坐的公交车数量。返回 -1 表示不可能到达终点车站。



示例：

输入：
routes = [[1, 2, 7], [3, 6, 7]]
S = 1
T = 6
输出：2
解释：
最优策略是先乘坐第一辆公交车到达车站 7, 然后换乘第二辆公交车到车站 6。

[[7,12],[4,5,15],[6],[15,19],[9,12,13]]
15
12

-1

提示：

1 <= routes.length <= 500.
1 <= routes[i].length <= 10^5.
0 <= routes[i][j] < 10 ^ 6.

分析：
1. 给公交车编号，找出到达开始位置的公交车和达到终点位置的公交车
2. 使用二维数组，记录有相交车站的公交车
3. 使用广度优先遍历，查看有相交的公交车是否能够到达终点站
*/

var minR int = -1

func hasEdge(aRoute []int, bRoute []int) bool {
	for i := 0; i < len(aRoute); i++ {
		for j := 0; j < len(bRoute); j++ {
			if aRoute[i] == bRoute[j] {
				return true
			}
		}
	}
	return false
}

func NumBusesToDestination(routes [][]int, S int, T int) int {
	minR = -1
	length := len(routes)
	if length == 0 {
		return -1
	}
	if S == T {
		return 0
	}
	//sb和eb存储经过开始和终点的公交车编号
	var sb, eb []int
	for i := 0; i < length; i++ {
		for j := 0; j < len(routes[i]); j++ {
			if routes[i][j] == S {
				sb = append(sb, i)
			}
			if routes[i][j] == T {
				eb = append(eb, i)
			}
		}
	}
	//判断两个公交车是否有相交的车站
	connect := make([][]int, length)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if hasEdge(routes[i], routes[j]) {
				connect[i] = append(connect[i], j)
				connect[j] = append(connect[j], i)
			}
		}
	}
	//fmt.Println(sb,eb,connect)
	//方案1：深度优先遍历.从有起始车站的公交车开始遍历
	/*for i := 0; i < len(sb); i++ {
		r := make([]int, length)
		r[sb[i]] = 1
		dfsRoute(connect, sb[i], eb, r, 1)
		r[sb[i]] = 0
	}*/

	//方案2：广度优先遍历
	for i := 0; i < len(sb); i++ {
		q := make([]int, 0)
		q = append(q, sb[i])
		bfsRoute(connect, q, eb, length)
	}
	fmt.Println(minR)
	return minR
}

func inArray(data int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if data == arr[i] {
			return true
		}
	}
	return false
}

//广度优先搜索
func bfsRoute(connect [][]int, q []int, endBusList []int, length int) {
	if inArray(q[0], endBusList) {
		minR = 1
		return
	}
	record := make([]int, length)
	record[q[0]] = 1
	for i := 0; i < length && i < len(q); i++ {
		startBus := q[i]
		for j := 0; j < len(connect[startBus]); j++ {
			if record[connect[startBus][j]] == 0 {
				q = append(q, connect[startBus][j])
				record[connect[startBus][j]] = record[startBus] + 1
				if inArray(connect[startBus][j], endBusList) {
					if minR == -1 || minR > record[startBus]+1 {
						minR = record[startBus] + 1
					}
					return
				}
			}
		}
	}
}

//深度优先遍历
func dfsRoute(connect [][]int, startBus int, endBusList []int, r []int, l int) {
	if minR != -1 && l >= minR {
		return
	}
	if inArray(startBus, endBusList) {
		if minR == -1 || minR > l {
			minR = l
		}
		return
	}
	for i := 0; i < len(connect[startBus]); i++ {
		busIndex := connect[startBus][i]
		//fmt.Println(busIndex,r,endBusList)
		if r[busIndex] == 0 {
			if inArray(busIndex, endBusList) {
				if minR == -1 || minR > l+1 {
					minR = l + 1
				}
			} else {
				r[busIndex] = l + 1
				dfsRoute(connect, busIndex, endBusList, r, l+1)
				r[busIndex] = 0
			}
		}
	}
}
