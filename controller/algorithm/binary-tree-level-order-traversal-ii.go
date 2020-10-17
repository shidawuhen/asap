package algorithm

import "container/list"

/*
原题：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
107. 二叉树的层次遍历 II
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/

type NodeLevel struct {
	node  *TreeNode
	level int
}

func levelOrderBottom(root *TreeNode) [][]int {
	var reslevel [][]int
	if root == nil {
		return reslevel
	}
	record := make(map[int][]int)
	queue := list.New()
	queue.PushBack(NodeLevel{root, 0})
	e := queue.Front()
	for e != nil {
		//获取头部
		v := e.Value.(NodeLevel)
		if v.node != nil {
			queue.PushBack(NodeLevel{v.node.Left, v.level + 1})
			queue.PushBack(NodeLevel{v.node.Right, v.level + 1})
			if _, ok := record[v.level]; !ok {
				record[v.level] = make([]int, 0)
			}
			record[v.level] = append(record[v.level], v.node.Val)
		}
		queue.Remove(e)
		e = queue.Front()
	}
	lenght := len(record)
	reslevel = make([][]int, lenght)
	for level, arr := range record {
		reslevel[lenght-level-1] = arr
	}
	return reslevel
}
