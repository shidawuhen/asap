package algorithm

/*
原题：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
剑指 Offer 32 - III. 从上到下打印二叉树 III
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]


提示：

节点总数 <= 1000
*/

func levelOrderWithLevelReverse(root *TreeNode) [][]int {
	levelTree := make([][]int, 0)
	if root == nil {
		return levelTree
	}
	queue := make([]*TreeNode, 0)
	levelRecord := make([]int, 0)
	queue = append(queue, root)
	levelRecord = append(levelRecord, 0)
	for i := 0; i < len(queue); i++ {
		if len(levelTree) <= levelRecord[i] {
			levelTree = append(levelTree, []int{})
		}
		levelTree[levelRecord[i]] = append(levelTree[levelRecord[i]], queue[i].Val)
		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
			levelRecord = append(levelRecord, levelRecord[i]+1)
		}
		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
			levelRecord = append(levelRecord, levelRecord[i]+1)
		}

	}
	for i := 0; i < len(levelTree); i++ {
		if i%2 != 0 {
			for j := 0; j < len(levelTree[i])/2; j++ {
				levelTree[i][j], levelTree[i][len(levelTree[i])-1-j] = levelTree[i][len(levelTree[i])-1-j], levelTree[i][j]
			}
		}
	}
	return levelTree
}
