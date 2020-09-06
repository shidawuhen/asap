package algorithm

/*
原题：https://leetcode-cn.com/problems/minimum-absolute-difference/
1200. 最小绝对差
给你个整数数组 arr，其中每个元素都 不相同。

请你找到所有具有最小绝对差的元素对，并且按升序的顺序返回。

示例 1：

输入：arr = [4,2,1,3]
输出：[[1,2],[2,3],[3,4]]
示例 2：

输入：arr = [1,3,6,10,15]
输出：[[1,3]]
示例 3：

输入：arr = [3,8,-10,23,19,-4,-14,27]
输出：[[-14,-10],[19,23],[23,27]]


提示：
2 <= arr.length <= 10^5
-10^6 <= arr[i] <= 10^6

分析：
使用二分法计算，先使用快速排序将数组排序，然后用二分法计算出最短距离，最后遍历数组，找到所有距离为最小值的数值对
*/

func MinimumAbsDifference(arr []int) [][]int {
	res := make([][]int, 0)
	length := len(arr)
	if length < 2 {
		return res
	}
	//排序
	quickSort(arr, 0, length-1)
	//先计算出最小值
	minV := computeDistance(arr, 0, length-1)
	//遍历数组，判断距离为minV的数对
	for i := 1; i < length; i++ {
		if arr[i]-arr[i-1] == minV {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}

//迭代计算距离
func computeDistance(arr []int, startIndex int, endIndex int) int {
	maxV := 10000000
	if startIndex >= endIndex {
		return maxV
	}
	dv := []int{maxV, maxV, maxV, maxV}
	middleIndex := (startIndex + endIndex) / 2
	dv[0] = computeDistance(arr, startIndex, middleIndex-1)
	dv[1] = computeDistance(arr, middleIndex+1, endIndex)
	if middleIndex-1 >= startIndex && arr[middleIndex]-arr[middleIndex-1] < dv[2] {
		dv[2] = arr[middleIndex] - arr[middleIndex-1]
	}
	if middleIndex+1 <= endIndex && arr[middleIndex+1]-arr[middleIndex] < dv[3] {
		dv[3] = arr[middleIndex+1] - arr[middleIndex]
	}
	d := maxV
	for _, v := range dv {
		if d > v {
			d = v
		}
	}
	return d
}
