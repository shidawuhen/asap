package algorithm

/*
原题：https://leetcode-cn.com/problems/smallest-k-lcci/
面试题 17.14. 最小K个数
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

示例：

输入： arr = [1,3,5,7,2,4,6,8], k = 4
输出： [1,2,3,4]
提示：

0 <= len(arr) <= 100000
0 <= k <= min(100000, len(arr))

分析：
使用堆来完成
1. 使用插入法将数组调整为最小堆结构
2. 将根节点和最大值互换，然后使用筛选调整法找到次小的值

*/

func SmallestK(arr []int, k int) []int {
	length := len(arr)
	if length == 0 {
		return []int{}
	}
	//1.调整为最小堆
	for i := length/2 - 1; i >= 0; i-- {
		siftHeap(arr, i)
	}
	//2.将第一个值和最后一个值互换，然后使用筛选法将剩余的数据重新调整为堆
	//fmt.Println(arr)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[0]
		arr[0], arr[length-1-i] = arr[length-1-i], arr[0]
		siftHeap(arr[0:length-1-i], 0)
		//fmt.Println(arr)
	}
	//fmt.Println(arr,res)
	return res
}

//arr是数组，index表示要处理的节点
func siftHeap(arr []int, index int) {
	length := len(arr)
	j := 2*index + 1
	for j < length {
		//最小堆需要找最小的值，比较父节点和左右子节点，找出最小的值
		if j+1 < length && (arr[j] > arr[j+1]) {
			j++
		}
		if arr[index] > arr[j] {
			arr[index], arr[j] = arr[j], arr[index]
			index = j
			j = 2*index + 1
		} else {
			break
		}
	}
}
