package algorithm

/*
原题：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
剑指 Offer 34. 二叉树中和为某一值的路径
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。



示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]


提示：

节点总数 <= 10000

分析：深度优先遍历

*/
var recordPath [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	recordPath = make([][]int, 0)
	if root == nil {
		return recordPath
	}
	path := make([]int, 0)
	path = append(path, root.Val)
	sum -= root.Val
	pathSumRecord(root, sum, path)
	return recordPath
}

func pathSumRecord(root *TreeNode, sum int, path []int) {
	//fmt.Println(path,sum,root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == 0 { //成功了
			temp := make([]int, 0)
			temp = append(temp, path...)
			recordPath = append(recordPath, temp)
			return
		} else {
			return
		}
	}
	if root.Left != nil {
		path = append(path, root.Left.Val)
		pathSumRecord(root.Left, sum-root.Left.Val, path)
		path = path[0 : len(path)-1]
	}
	if root.Right != nil {
		path = append(path, root.Right.Val)
		pathSumRecord(root.Right, sum-root.Right.Val, path)
		path = path[0 : len(path)-1]
	}
}
