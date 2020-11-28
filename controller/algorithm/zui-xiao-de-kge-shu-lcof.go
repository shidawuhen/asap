package algorithm

import (
	"container/heap"
)

/*
原题：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
剑指 Offer 40. 最小的k个数
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。



示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]


限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000

分析
1. 先将数组使用最小堆排序
2. 将最小值和尾部值替换，然后重新排序，当然排序位置截止到交换的位置

*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func GetLeastNumbers(arr []int, k int) []int {
	res := []int{}
	length := len(arr)
	if length == 0 {
		return res
	}
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < length; i++ {
		heap.Push(h, arr[i])
	}
	for i := 0; i < k; i++ {
		res = append(res, heap.Remove(h, 0).(int))
	}
	return res
}
