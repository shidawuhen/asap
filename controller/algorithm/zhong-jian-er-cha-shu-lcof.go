package algorithm

/*
原题：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
剑指 Offer 07. 重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


限制：

0 <= 节点个数 <= 5000



注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

分析：
1.使用递归，根据前序遍历的点，切割中序遍历
2.另一个核心点是切割中序遍历后，左右的长度用来切割前序遍历集合，因为前序和中序的左子树和右子树成团存在
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	return createTree(preorder, inorder)
}

func createTree(preorder []int, inorder []int) *TreeNode {
	plength := len(preorder)
	if plength == 0 {
		return nil
	}
	if plength == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}

	root := &TreeNode{}

	ilength := len(inorder)
	for i := 0; i < ilength; i++ {
		if inorder[i] == preorder[0] { //如果前序遍历的点在中序遍历里找到了对应点
			root.Val = preorder[0]
			root.Left = createTree(preorder[1:i+1], inorder[0:i])
			root.Right = createTree(preorder[i+1:plength], inorder[i+1:ilength])
			break
		}
	}

	return root
}
