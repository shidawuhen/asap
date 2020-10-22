package algorithm

import (
	"container/heap"
	"sort"
)

/*
原题：https://leetcode-cn.com/problems/maximum-performance-of-a-team/
1383. 最大的团队表现值
公司有编号为 1 到 n 的 n 个工程师，给你两个数组 speed 和 efficiency ，
其中 speed[i] 和 efficiency[i] 分别代表第 i 位工程师的速度和效率。请你返回由最多 k 个工程师组成的 ​​​​​​最大团队表现值 ，
由于答案可能很大，请你返回结果对 10^9 + 7 取余后的结果。

团队表现值 的定义为：一个团队中「所有工程师速度的和」乘以他们「效率值中的最小值」。



示例 1：

输入：n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
输出：60
解释：
我们选择工程师 2（speed=10 且 efficiency=4）和工程师 5（speed=5 且 efficiency=7）。他们的团队表现值为 performance = (10 + 5) * min(4, 7) = 60 。
示例 2：

输入：n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 3
输出：68
解释：
此示例与第一个示例相同，除了 k = 3 。我们可以选择工程师 1 ，工程师 2 和工程师 5 得到最大的团队表现值。表现值为 performance = (2 + 10 + 5) * min(5, 4, 7) = 68 。
示例 3：

输入：n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 4
输出：72


提示：

1 <= n <= 10^5
speed.length == n
efficiency.length == n
1 <= speed[i] <= 10^5
1 <= efficiency[i] <= 10^8
1 <= k <= n

分析：
题目要求我们最优化「速度和」和「效率最小值」的乘积。变化的量有两个，一个是「速度」，一个是「效率」，这看起来有些棘手。我们不妨采用「动一个，定一个」的策略——即我们可以枚举效率的最小值 e_{\min}e
min
​
 ，在所有效率大于 e_{\min}e
min
​
  的工程师中选取不超过 k - 1k−1 个，让他们的速度和最大。

思考：为什么是 k - 1k−1 个而不是 kk 个？ 因为最小值 e_{\min}e
min
​
  代表的工程师是必选，加起来一共 kk 个，所以剩下只要选 k - 1k−1 个。

思考：如何满足速度和最大？ 因为 speed[i] > 0，所以只需要选效率大于 e_{\min}e
min
​
  中速度最大的 k - 1k−1 个，如果效率大于 e_{\min}e
min
​
  的元素小于 k - 1k−1，就全取。

具体地，我们可以对工程师先按照「效率」从高到低的规则排序，然后从前往后枚举这个序列中的元素作为 e_{\min}e
min
​
 ，这样可以保证前面的元素的效率都比当前这个工程师高，然后维护一个以「速度」为关键字的小根堆，存放前面已经枚举过的元素中速度前 k - 1k−1 大的，动态维护这个堆的速度和，一轮枚举后，我们可以得到乘积最大值。

思考：为什么是小根堆？ 因为我们要动态维护前 k - 1k−1 大的元素，当堆内的元素超过 k - 1k−1 的时候，我们可以从堆顶 pop 掉比较小的元素，保证最大的 k - 1k−1 个元素还在堆中。

*/

func MaxPerformance(n int, speed []int, efficiency []int, k int) int {
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return efficiency[indexes[i]] > efficiency[indexes[j]]
	})
	ph := speedHeap{}
	heap.Init(&ph)
	var speedSum int
	var max int64
	for _, index := range indexes {
		if ph.Len() == k { //因为每一个经过的speed我都加了，所以当超出k的时候，需要把最小的去掉。因为effective相同，speed越大越好
			speedSum -= heap.Pop(&ph).(int)
		}
		speedSum += speed[index]
		heap.Push(&ph, speed[index])
		max = Max(max, int64(speedSum)*int64(efficiency[index]))
	}
	return int(max % (1e9 + 7))
}

type speedHeap []int

func (h speedHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h speedHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h speedHeap) Len() int            { return len(h) }
func (h *speedHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *speedHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:h.Len()-1]
	return res
}
func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
