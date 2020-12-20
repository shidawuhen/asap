package algorithm

/*
原题：
https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
剑指 Offer 68 - II. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]





示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

分析：
一种方法是分别计算p和q的路径，然后根据路径判断
另一种是判断返回的左右子树和父亲
*/

func lowestCommonAncestorNew(root, p, q *TreeNode) *TreeNode {
	_, _, parent := findAnc(root, p, q)
	return parent
}

func findAnc(root, p, q *TreeNode) (l, r, parent *TreeNode) {
	if root == nil {
		return nil, nil, nil
	}

	l0, r0, parent := findAnc(root.Left, p, q)
	if parent != nil {
		return l0, r0, parent
	}

	l1, r1, parent := findAnc(root.Right, p, q)
	if parent != nil {
		return l1, r1, parent
	}

	//判断左侧子树和右侧子树是否有所有的p q
	pexist := isSameData(l0, r0, l1, r1, root, p)
	qexist := isSameData(l0, r0, l1, r1, root, q)
	if pexist != nil && qexist != nil {
		return pexist, qexist, root
	}
	return pexist, qexist, nil

}

func isSameData(l0, r0, l1, r1, root, node *TreeNode) *TreeNode {
	if l0 != nil && l0.Val == node.Val {
		return l0
	}
	if r0 != nil && r0.Val == node.Val {
		return r0
	}
	if l1 != nil && l1.Val == node.Val {
		return l1
	}
	if r1 != nil && r1.Val == node.Val {
		return r1
	}
	if root != nil && root.Val == node.Val {
		return root
	}
	return nil
}
