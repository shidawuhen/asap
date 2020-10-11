package algorithm

import (
	"strconv"
)

/*
原题：https://leetcode-cn.com/problems/binary-tree-paths/
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

分析：
这个一道很基础的深度遍历习题，值得掌握
*/

/* Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	path := ""
	res = getPath(root, path)
	//fmt.Println(res)
	return res
}
func getPath(node *TreeNode, path string) (res []string) {
	if node == nil {
		return
	}
	//根节点
	if node.Left == nil && node.Right == nil {
		path += strconv.Itoa(node.Val)
		res = append(res, path)
		return
	}
	if node.Left != nil {
		r := getPath(node.Left, path+strconv.Itoa(node.Val)+"->")
		res = append(res, r...)
	}
	if node.Right != nil {
		r := getPath(node.Right, path+strconv.Itoa(node.Val)+"->")
		res = append(res, r...)
	}
	return res
}
