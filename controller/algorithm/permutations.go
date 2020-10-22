package algorithm

/*
原题：https://leetcode-cn.com/problems/permutations
给定一个 没有重复 数字的序列，返回其所有可能的全排列。
示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

分析：
当前n-1个数字排序后，会有(n-1)!的数组，需要将第n个数字放到每一个数组的每一个可能位置
*/

func Permute(nums []int) [][]int {
	length := len(nums)
	totalArray := 1
	oldA := make([][]int, 1)
	for n := 0; n < 1; n++ {
		oldA[n] = make([]int, 1)
	}
	oldA[0][0] = nums[0]
	for i := 1; i < length; i++ { //代表n，一直到最后一个值
		nextTotalArray := totalArray * (i + 1) //(n)!
		newA := make([][]int, nextTotalArray)
		for n := 0; n < nextTotalArray; n++ {
			newA[n] = make([]int, i+1)
		}
		for j := 0; j < totalArray; j++ { //代表(n-1)!里的每一个数组
			for k := 0; k < i+1; k++ { //代表每一个数组里的间隔位置，最头和最尾也需要计算上，将nums[i]的值放到间隔上
				//重新放入新位置
				newA[j*(i+1)+k][k] = nums[i]
				for t := 0; t < len(oldA[j]); t++ {
					if t < k {
						newA[j*(i+1)+k][t] = oldA[j][t]
					} else {
						newA[j*(i+1)+k][t+1] = oldA[j][t]
					}
				}
			}
		}
		oldA = newA
		totalArray = nextTotalArray
	}
	return oldA
}
