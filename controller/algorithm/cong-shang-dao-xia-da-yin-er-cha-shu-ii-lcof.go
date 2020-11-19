package algorithm

/*
原题：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
剑指 Offer 32 - II. 从上到下打印二叉树 II
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。



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
  [9,20],
  [15,7]
]


提示：

节点总数 <= 1000
*/

func levelOrderWithLevel(root *TreeNode) [][]int {
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
	return levelTree
}
