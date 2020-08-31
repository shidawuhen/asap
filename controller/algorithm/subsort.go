package algorithm

/*
原题：https://leetcode-cn.com/problems/sub-sort-lcci/

给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。

示例：

输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
输出： [3,9]
提示：

0 <= len(array) <= 1000000

分析：
解法1：
练习快速排序，将排序后的值与原序列的值对比，最前面变化的位置和最后面变化的位置便是变更位置。结果是超时，但是算法是对的
*/

func partition(array []int, startIndex int, endIndex int) int {
	i := startIndex
	j := endIndex
	for i < j {
		for i < j && array[i] <= array[j] {
			j--
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
			i++
		}
		for i < j && array[i] <= array[j] {
			i++
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
			j--
		}
	}
	return i
}

func quickSort(array []int, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	middleIndex := partition(array, startIndex, endIndex)
	quickSort(array, startIndex, middleIndex)
	quickSort(array, middleIndex+1, endIndex)
}

func SubSort(array []int) []int {
	res := []int{-1, -1}
	if len(array) == 0 {
		return res
	}
	length := len(array)
	oldArray := make([]int, length)
	for i := 0; i < length; i++ {
		oldArray[i] = array[i]
	}
	quickSort(array, 0, len(array)-1)

	for i := 0; i < length-1; i++ {
		if array[i] != oldArray[i] {
			res[0] = i
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		if array[i] != oldArray[i] {
			res[1] = i
			break
		}
	}
	if res[0] >= res[1] {
		return []int{-1, -1}
	}

	return res
}
