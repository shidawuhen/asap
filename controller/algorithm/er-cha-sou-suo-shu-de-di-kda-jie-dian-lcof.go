package algorithm

/*
原题：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
剑指 Offer 54. 二叉搜索树的第k大节点
给定一棵二叉搜索树，请找出其中第k大的节点。



示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4


限制：

1 ≤ k ≤ 二叉搜索树元素个数

分析：先遍历右子树，然后中间节点，然后左子树
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var realK int
var recordNode int

func kthLargest(root *TreeNode, k int) int {
	realK = k
	recordNode = 0
	findK(root)
	return recordNode
}

func findK(node *TreeNode) {
	if node == nil {
		return
	}

	findK(node.Right)
	realK--
	if realK == 0 {
		recordNode = node.Val
		return
	}
	//fmt.Println(node.Val)
	findK(node.Left)
}
