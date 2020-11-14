package algorithm

/*
原题：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
剑指 Offer 26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true
限制：

0 <= 节点个数 <= 10000

*/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	return compareTree(A, B, false)
}

func compareTree(tree *TreeNode, B *TreeNode, startCompare bool) bool {
	if tree == nil || B == nil {
		if tree == nil && B != nil {
			return false
		}
		if tree == nil && B == nil && startCompare == true {
			return true
		}
		if tree == nil && B == nil && startCompare == false {
			return false
		}
		if tree != nil && B == nil && startCompare == true {
			return true
		}
		if tree != nil && B == nil && startCompare == false {
			return false
		}
	}
	if tree.Val == B.Val {
		res1 := compareTree(tree.Left, B.Left, true)
		res2 := compareTree(tree.Right, B.Right, true)
		if res1 == true && res2 == true {
			return true
		}
	}
	if startCompare == true {
		return false
	}
	res1 := compareTree(tree.Left, B, false)
	if res1 == true {
		return true
	}
	res2 := compareTree(tree.Right, B, false)
	if res2 == true {
		return true
	}
	return false
}
