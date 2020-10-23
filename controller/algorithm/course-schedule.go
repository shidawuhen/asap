package algorithm

/*
原题：https://leetcode-cn.com/problems/course-schedule/
207. 课程表
你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]

给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？



示例 1:

输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
示例 2:
          s  p
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。


提示：

输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
1 <= numCourses <= 10^5


分析：
这道题的核心思路是，不断将没有父亲的节点删除，与该父亲连接的子节点的父亲会减一，如果没有回路，则所有节点都可以删除，
否则一定无法删除所有节点

这个想法还是很巧妙的
*/

func CanFinish(numCourses int, prerequisites [][]int) bool {
	length := len(prerequisites)
	if length == 0 {
		return true
	}
	parentNum := make([]int, numCourses)
	parentSonList := make(map[int][]int)
	for i := 0; i < length; i++ {
		parentNum[prerequisites[i][0]]++
		parentSonList[prerequisites[i][1]] = append(parentSonList[prerequisites[i][1]], prerequisites[i][0])
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if parentNum[i] == 0 {
			queue = append(queue, i)
		}
	}

	for i := 0; i < len(queue); i++ {
		numCourses--
		index := queue[i]
		if sonList, ok := parentSonList[index]; ok {
			for j := 0; j < len(sonList); j++ {
				parentNum[sonList[j]]--
				if parentNum[sonList[j]] == 0 {
					queue = append(queue, sonList[j])
				}
			}
		}
	}
	if numCourses == 0 {
		return true
	} else {
		return false
	}
}
