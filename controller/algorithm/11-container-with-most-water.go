/**
@author: Jason Pang
@desc:
@date: 2021/12/4
**/
package algorithm

/**
 * 原题：https://leetcode-cn.com/problems/container-with-most-water/
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。

示例 1：



输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1
示例 3：

输入：height = [4,3,2,1,4]
输出：16
示例 4：

输入：height = [1,2,1]
输出：2


提示：

n == height.length
2 <= n <= 105
0 <= height[i] <= 104

思路：找每个高度之上，最左、最右的位置
*/
type PosWater struct {
	l int
	r int
}

func maxArea(height []int) int {
	//找出最大高度
	maxH := 0
	pos := make(map[int]PosWater, 0) //某个高度，最左或者最右的x
	for x, h := range height {
		if h > maxH {
			maxH = h
		}
		//判断该高度，最左最右在的位置x
		if p, ok := pos[h]; ok {
			if x < p.l {
				p.l = x
			}
			if x > p.r {
				p.r = x
			}
			pos[h] = p
		} else {
			pos[h] = PosWater{x, x}
		}
	}
	//fmt.Println(pos)
	//从高到低，计算pos 最左左右值
	for h := maxH - 1; h >= 0; h-- {
		if p, ok := pos[h]; ok {
			if pos[h].l >= pos[h+1].l {
				p.l = pos[h+1].l
			}
			if pos[h].r <= pos[h+1].r {
				p.r = pos[h+1].r
			}
			pos[h] = p
		} else {
			pos[h] = pos[h+1]
		}
	}
	//fmt.Println(pos)
	maxArea := 0
	for i := 0; i <= maxH; i++ { //遍历所有高度
		if (pos[i].r-pos[i].l)*i > maxArea {
			maxArea = (pos[i].r - pos[i].l) * i
		}
	}
	return maxArea
}
