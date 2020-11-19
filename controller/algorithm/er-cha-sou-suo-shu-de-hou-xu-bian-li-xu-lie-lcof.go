package algorithm

/*
原题：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
剑指 Offer 33. 二叉搜索树的后序遍历序列
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true


提示：

数组长度 <= 1000
*/

func verifyPostorder(postorder []int) bool {

	return isSearchTree(postorder)
}

func isSearchTree(postorder []int) bool {
	if len(postorder) == 0 || len(postorder) == 1 {
		return true
	}
	//判断最后一位，能够将其余数据切割成正确的两部分
	val := postorder[len(postorder)-1]
	splitIndex := -1
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] > val {
			splitIndex = i
			break
		}
	}
	//fmt.Println(postorder,splitIndex)
	if splitIndex != -1 {
		for i := splitIndex + 1; i < len(postorder)-1; i++ {
			if postorder[i] < val {
				return false
			}
		}
	}

	res1 := true
	if splitIndex != -1 {
		res1 = isSearchTree(postorder[0:splitIndex])
		res2 := isSearchTree(postorder[splitIndex : len(postorder)-1])
		return res1 && res2
	} else {
		return isSearchTree(postorder[0 : len(postorder)-1])
	}
}
