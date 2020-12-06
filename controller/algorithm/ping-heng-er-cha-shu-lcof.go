package algorithm

import "math"

/*
原题：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
剑指 Offer 55 - II. 平衡二叉树
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。



示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。



限制：

1 <= 树的结点个数 <= 10000
*/

func isBalanced(root *TreeNode) bool {
	_, b := isBalanceTree(root, 0)
	return b
}

func isBalanceTree(node *TreeNode, depth int) (d int, balance bool) {
	if node == nil {
		return depth, true
	}
	ld, lb := isBalanceTree(node.Left, depth+1)
	if lb == false {
		return -1, false
	}
	rd, rb := isBalanceTree(node.Right, depth+1)
	if rb == false {
		return -1, false
	}
	if math.Abs(float64(ld-rd)) <= 1 {
		return int(math.Max(float64(ld), float64(rd))), true
	}
	return -1, false
}
