package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/4sum/
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

*/

func FourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	if length < 4 {
		return res
	}
	rem := make(map[string]int)
	quickSort(nums, 0, length-1)
	for i := 0; i < length-3; i++ {
		for j := i + 1; j < length-2; j++ {
			for k := j + 1; k < length-1; k++ {
				for m := k + 1; m < length; m++ {
					if nums[i]+nums[j]+nums[k]+nums[m] == target {
						record := fmt.Sprintf("%s%s%s%s", nums[i], nums[j], nums[k], nums[m])
						if _, ok := rem[record]; !ok {
							res = append(res, []int{nums[i], nums[j], nums[k], nums[m]})
							rem[record] = 1
						}
					}
				}
			}
		}
	}
	//fmt.Println(res)
	return res
}
