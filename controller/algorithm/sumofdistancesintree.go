package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/sum-of-distances-in-tree/
834. 树中距离之和
给定一个无向、连通的树。树中有 N 个标记为 0...N-1 的节点以及 N-1 条边 。

第 i 条边连接节点 edges[i][0] 和 edges[i][1] 。

返回一个表示节点 i 与其他所有节点距离之和的列表 ans。

示例 1:

输入: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
输出: [8,12,6,10,10,10]
解释:
如下为给定的树的示意图：
  0
 / \
1   2
   /|\
  3 4 5

我们可以计算出 dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
也就是 1 + 1 + 2 + 2 + 2 = 8。 因此，answer[0] = 8，以此类推。
说明: 1 <= N <= 10000

[0 1 1 2 2 2]  8
[1 0 2 3 3 3]  12
[1 2 0 1 1 1]  6
[2 3 1 0 2 2]  10
[2 3 1 2 0 2] 10
[2 3 1 2 2 0] 10


 i 0 s 1  e 2
*/

//深度优先遍历
func distance(parentIndex int, startIndex int, endIndex int, linkMap map[int][]int, sum int) int {
	list, _ := linkMap[startIndex]
	for i := 0; i < len(list); i++ {
		if list[i] != parentIndex {
			if list[i] == endIndex {
				return sum + 1
			} else {
				dis := distance(startIndex, list[i], endIndex, linkMap, sum+1)
				if dis != 0 {
					return dis
				}
			}
		}
	}
	return 0
}

func SumOfDistancesInTree(N int, edges [][]int) []int {
	res := make([]int, N)
	//linkMap用于计算连接关系
	linkMap := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		linkMap[edges[i][0]] = append(linkMap[edges[i][0]], edges[i][1])
		linkMap[edges[i][1]] = append(linkMap[edges[i][1]], edges[i][0])
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if i != j {
				sum := distance(-1, i, j, linkMap, 0)
				res[i] += sum
				res[j] += sum
			}
		}
	}
	fmt.Println(res)
	return res
}
