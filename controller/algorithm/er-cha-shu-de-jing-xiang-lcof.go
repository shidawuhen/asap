package algorithm

/*
原题：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
剑指 Offer 27. 二叉树的镜像
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1



示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]


限制：

0 <= 节点个数 <= 1000
*/

func mirrorTree(root *TreeNode) *TreeNode {
	changePos(root)
	return root
}

func changePos(index *TreeNode) {
	if index == nil {
		return
	}
	changePos(index.Left)
	changePos(index.Right)
	index.Left, index.Right = index.Right, index.Left
}
