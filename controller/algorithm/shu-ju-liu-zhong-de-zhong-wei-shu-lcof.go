package algorithm

/*
原题：https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/
剑指 Offer 41. 数据流中的中位数
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例 1：

输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]
示例 2：

输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]


限制：

最多会对 addNum、findMedian 进行 50000 次调用。

分析：
应该用大根堆和小根堆的方案，但是我直接用二分了
*/

type MedianFinder struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		nums: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	//寻找第一个<=num的位置
	startIndex := 0
	endIndex := len(this.nums) - 1
	for startIndex <= endIndex {
		middleIndex := (startIndex + endIndex) / 2
		if this.nums[middleIndex] >= num {
			endIndex = middleIndex - 1
		} else {
			startIndex = middleIndex + 1
		}
	}
	temp := make([]int, 0)
	temp = append(temp, this.nums[0:startIndex]...)
	temp = append(temp, num)
	temp = append(temp, this.nums[startIndex:len(this.nums)]...)
	//this.nums = append(append(this.nums[0:startIndex],num),this.nums[startIndex:len(this.nums)]...)
	this.nums = temp
	//fmt.Println(startIndex,num,this.nums)
}

func (this *MedianFinder) FindMedian() float64 {
	length := len(this.nums)
	if length%2 == 0 { //偶数
		return (float64(this.nums[length/2]) + float64(this.nums[(length-1)/2])) / 2
	} else {
		return float64(this.nums[length/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
