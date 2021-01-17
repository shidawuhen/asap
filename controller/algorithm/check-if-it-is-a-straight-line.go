package algorithm

import "math"

/*
原题：https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/
1232. 缀点成线
在一个 XY 坐标系中有一些点，我们用数组 coordinates 来分别记录它们的坐标，其中 coordinates[i] = [x, y] 表示横坐标为 x、纵坐标为 y 的点。

请你来判断，这些点是否在该坐标系中属于同一条直线上，是则返回 true，否则请返回 false。



示例 1：



输入：coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
输出：true
示例 2：



输入：coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
输出：false


提示：

2 <= coordinates.length <= 1000
coordinates[i].length == 2
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
coordinates 中不含重复的点
*/

func checkStraightLine(coordinates [][]int) bool {
	length := len(coordinates)
	if length <= 2 {
		return true
	}
	//直线公式
	//（y-y1）/(y2-y1)＝(x-x1)/(x2-x1)
	p0 := coordinates[0]
	p1 := coordinates[1]
	x1 := p0[0]
	y1 := p0[1]
	x2 := p1[0]
	y2 := p1[1]
	for i := 2; i < length; i++ {
		x := coordinates[i][0]
		y := coordinates[i][1]
		if math.Abs(float64(y-y1)/float64(y2-y1)-float64(x-x1)/float64(x2-x1)) > 0.000001 {
			return false
		}
	}
	return true

}
