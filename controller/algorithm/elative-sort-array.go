package algorithm

/*
原题：https://leetcode-cn.com/problems/relative-sort-array/
1122. 数组的相对排序
给你两个数组，arr1 和 arr2，

arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。



示例：

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]


提示：

arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
arr2 中的元素 arr2[i] 各不相同
arr2 中的每个元素 arr2[i] 都出现在 arr1 中
*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	record := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		record[arr1[i]]++
	}
	notExists := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(arr2); i++ {
		length, _ := record[arr2[i]]
		for j := 0; j < length; j++ {
			res = append(res, arr2[i])
		}
		delete(record, arr2[i])
	}
	for v, num := range record {
		for j := 0; j < num; j++ {
			notExists = append(notExists, v)
		}
	}
	quickSort(notExists, 0, len(notExists)-1)
	res = append(res, notExists...)
	return res
}
