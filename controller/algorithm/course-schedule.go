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

输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。


提示：

输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
1 <= numCourses <= 10^5
*/


func CanFinish(numCourses int, prerequisites [][]int) bool {
	length := len(prerequisites)
	if length == 0 {
		return true
	}
	recordMap := make(map[int][]int)
	for i := 0; i < length; i++ {
		recordMap[prerequisites[i][0]] = append(recordMap[prerequisites[i][0]],prerequisites[i][1:len(prerequisites[i])]...)
	}
	//fmt.Println(recordMap)
	for i := 0; i < length; i++ {

		queue := make([]int,1)
		queue[0] = prerequisites[i][0]
		used := make([]bool,numCourses)
		used[prerequisites[i][0]] = true
		//fmt.Println(queue)
		for j := 0;j < len(queue);j++  {
			//fmt.Println(j,queue,used)
			index := queue[j]
			if list, ok:= recordMap[index];ok{
				for m:=0;m < len(list);m++{
					if used[list[m]] == true{
						return false
					}else{
						used[list[m]] = true
					}
				}
				queue = append(queue,list...)
			}
		}

	}
	return true
}