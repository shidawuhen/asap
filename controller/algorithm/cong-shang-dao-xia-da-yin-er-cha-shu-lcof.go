package algorithm

/*
原题：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
剑指 Offer 32 - I. 从上到下打印二叉树
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]


提示：

节点总数 <= 1000
*/

func levelOrder(root *TreeNode) []int {
	levelTree := make([]int, 0)
	if root == nil {
		return levelTree
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for i := 0; i < len(queue); i++ {
		levelTree = append(levelTree, queue[i].Val)
		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
		}
		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
		}
	}
	return levelTree
}
