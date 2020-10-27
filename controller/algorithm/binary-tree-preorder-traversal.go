package algorithm

/*
原题：https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
144. 二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

var record []int

func preorderTraversal(root *TreeNode) []int {
	record = make([]int, 0)
	PrintT(root)
	return record
}

func PrintT(index *TreeNode) {
	if index != nil {
		record = append(record, index.Val)
		PrintT(index.Left)
		PrintT(index.Right)
	}
}
