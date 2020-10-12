package algorithm

import (
	"fmt"
	"math"
)

/*
原题：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
530. 二叉搜索树的最小绝对差
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

示例：

输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。


提示：

树中至少有 2 个节点。
本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同

*/

/*Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	data := getMin(root)
	minV := data[1] - data[0]
	for i := 1; i <= len(data)-2; i++ {
		if data[i+1]-data[i] < minV {
			minV = data[i+1] - data[i]
		}
	}
	fmt.Println(data, minV)
	return minV
}

func getMin(node *TreeNode) (data []int) {
	if node.Left != nil {
		res := getMin(node.Left)
		data = append(data, res...)
	}
	data = append(data, node.Val)
	//fmt.Println(node.Val)
	if node.Right != nil {
		res := getMin(node.Right)
		data = append(data, res...)
	}
	return
}
