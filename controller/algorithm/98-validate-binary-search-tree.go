/**
@author: Jason Pang
@desc:
@date: 2022/2/20
**/
package algorithm

/*
原题：https://leetcode-cn.com/problems/validate-binary-search-tree/
98. 验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。


示例 1：


输入：root = [2,1,3]
输出：true
示例 2：


输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。


提示：

树中节点数目范围在[1, 104] 内
-231 <= Node.val <= 231 - 1

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//错误解法
func isValidBSTWrong(root *TreeNode) bool {

	return dspValidBST(root)
}

func dspValidBST(node *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left != nil && node.Left.Val >= node.Val {
		return false
	}
	if node.Right != nil && node.Right.Val <= node.Val {
		return false
	}
	return dspValidBST(node.Left) && dspValidBST(node.Left)
}

//错误
//左子树确认最大值，右子树确认最小值
func isValidBSTWrong2(root *TreeNode) bool {
	if root == nil {
		return false
	}

	res, _ := newdspValidBST(root, "")
	return res
}

func newdspValidBST(node *TreeNode, dir string) (res bool, val int) {
	lIsVal, rIsVal := true, true
	lV, rV := node.Val, node.Val
	if node.Left != nil {
		lIsVal, lV = newdspValidBST(node.Left, "l")
	}
	if node.Right != nil {
		rIsVal, rV = newdspValidBST(node.Right, "r")
	}
	fmt.Println(node.Val, lIsVal, lV, rIsVal, rV)
	//如果有一个不合格，都不合格
	if lIsVal == false || rIsVal == false {
		return false, 0
	}
	//判断左右子树是否合规
	if (node.Left != nil && node.Val <= lV) || (node.Right != nil && node.Val >= rV) {
		return false, 0
	}
	if (node.Left != nil && node.Val <= node.Left.Val) || (node.Right != nil && node.Val >= node.Right.Val) {
		return false, 0
	}
	//判断极值
	if dir == "l" { //左子树确认最大值
		if node.Right != nil {
			return true, rV
		}
		return true, node.Val
	}
	if dir == "r" { //右子树确认最小值
		if node.Left != nil {
			return true, lV
		}
		return true, node.Val
	}
	return true, node.Val
}

//中序遍历，判断是否递增
var bstVal []int

func isValidBST(root *TreeNode) bool {
	bstVal = make([]int, 0)
	dspVal(root)
	//fmt.Println(bstVal)
	for i := 1; i < len(bstVal); i++ {
		if bstVal[i]-bstVal[i-1] <= 0 {
			return false
		}
	}
	return true
}

func dspVal(node *TreeNode) {
	if node == nil {
		return
	}
	dspVal(node.Left)
	bstVal = append(bstVal, node.Val)
	dspVal(node.Right)
}
