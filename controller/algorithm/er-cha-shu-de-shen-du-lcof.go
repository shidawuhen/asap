package algorithm
/*
原题：https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
剑指 Offer 55 - I. 二叉树的深度
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。



提示：

节点总数 <= 10000
*/
var maxTreeDepth int = 0
func maxDepth(root *TreeNode) int {
	maxTreeDepth = 0
	treeDepth(root,0)
	return maxTreeDepth
}

func treeDepth(node *TreeNode,depth int){
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if depth + 1 > maxTreeDepth{
			maxTreeDepth = depth + 1
		}
		return
	}
	treeDepth(node.Left,depth+1)
	treeDepth(node.Right,depth+1)
}