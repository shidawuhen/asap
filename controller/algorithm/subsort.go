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
解法2：
1.从左向右遍历，找到第一个不是继续增大的值，记录索引l，然后从右往左遍历，找到第一个不是继续减小的值，记录索引r
2.在l和r中找最大和最小值
3.从左向右遍历，找到第一个大于最小值的位置为左起点。从右往左遍历，找到第一个小于最大值的位置为右起点

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

//快速排序
func SubSort(array []int) []int {
	res := []int{-1, -1}
	if len(array) == 0 {
		return res
	}
	length := len(array)
	oldArray := make([]int, length)
	copy(oldArray,array)
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

//简单判断
func SubSortSimple(array []int) []int {
	leftIndex := -1
	rightIndex := -1
	res := []int{leftIndex, rightIndex}
	if len(array) == 0 {
		return res
	}
	length := len(array)
	minValue := -1
	lSpecialIndex := -1
	maxValue := -1
	rSpecialIndex := -1
	//找到左侧第一个不符合的位置
	for i := 0; i < length-1; i++ {
		if array[i] > array[i+1] {
			lSpecialIndex = i + 1
			break
		}
	}
	//说明肯定不存在
	if lSpecialIndex == -1 {
		return res
	}
	//找到右侧第一个不符合的位置
	for j := length - 1; j > 0; j-- {
		if array[j-1] > array[j] {
			rSpecialIndex = j - 1
			break
		}
	}
	//找两个index中的最小值和最大值
	minValue = array[lSpecialIndex]
	maxValue = array[lSpecialIndex]
	for i := lSpecialIndex + 1; i <= rSpecialIndex; i++ {
		if minValue > array[i] {
			minValue = array[i]
		}
		if maxValue < array[i] {
			maxValue = array[i]
		}
	}

	//判断左侧比两个special都大的第一个位置
	for i := 0; i <= lSpecialIndex; i++ {
		if array[i] > minValue {
			leftIndex = i
			break
		}
	}
	//判断右侧比两个special都小的第一个位置
	for j := length - 1; j >= rSpecialIndex; j-- {
		if array[j] < maxValue {
			rightIndex = j
			break
		}
	}
	return []int{leftIndex, rightIndex}
}
