package algorithm

/*
原题：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
剑指 Offer 28. 对称的二叉树
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3



示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false


限制：

0 <= 节点个数 <= 1000
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, root.Right)
}

func isSameTree(l *TreeNode, r *TreeNode) bool {
	if l == nil && r != nil {
		return false
	}
	if l != nil && r == nil {
		return false
	}
	if l == nil && r == nil {
		return true
	}
	if l.Val == r.Val {
		res1 := isSameTree(l.Left, r.Right)
		if res1 == false {
			return false
		}
		res2 := isSameTree(l.Right, r.Left)
		if res2 == false {
			return false
		}
		return true
	} else {
		return false
	}
}
