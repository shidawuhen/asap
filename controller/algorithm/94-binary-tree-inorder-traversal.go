/**
@author: Jason Pang
@desc:
@date: 2022/1/24
**/
package algorithm

/*
原题：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
94. 二叉树的中序遍历
给定一个二叉树的根节点 root ，返回它的 中序 遍历。



示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：


输入：root = [1,2]
输出：[2,1]
示例 5：


输入：root = [1,null,2]
输出：[1,2]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

思路：确定模式，然后写代码
*/

var nodeValue []int

func inorderTraversal(root *TreeNode) []int {
	nodeValue = make([]int, 0)
	if root == nil {
		return nodeValue
	}
	dsp(root)
	return nodeValue
}

func dsp(r *TreeNode) {
	if r.Left != nil {
		dsp(r.Left)
	}
	nodeValue = append(nodeValue, r.Val)
	if r.Right != nil {
		dsp(r.Right)
	}
}
