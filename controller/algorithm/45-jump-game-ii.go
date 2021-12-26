/**
@author: Jason Pang
@desc:
@date: 2021/12/22
**/
package algorithm

//方案1：动态规划
var recordJump map[int]int

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	recordJump = make(map[int]int, 0)
	recordJump[len(nums)-1] = 0
	return countJump(0, nums)
}

func countJump(index int, nums []int) int { //从index到末尾最短路径
	if index >= len(nums)-1 {
		return 0
	}
	minJump := 100000

	for i := 0; i < nums[index] && i+1+index < len(nums); i++ {
		j := 0
		if _, ok := recordJump[i+1+index]; ok {
			j = recordJump[i+1+index]
		} else {
			j = countJump(i+1+index, nums)
		}
		j = 1 + j
		if j < minJump {
			minJump = j
		}
	}

	recordJump[index] = minJump
	return minJump
}

//方案2 - 从后往前
func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	recordJump = make(map[int]int, 0)
	recordJump[len(nums)-1] = 0
	for index := len(nums) - 2; index >= 0; index-- {
		//根据自身能跳跃的长度，找到最小值
		minJump := 1000000
		for i := index + 1; i < index+1+nums[index] && i < len(nums); i++ {
			jump := 1 + recordJump[i]
			if jump < minJump {
				minJump = jump
			}
		}
		recordJump[index] = minJump
	}
	return recordJump[0]
}

//方案3 - 贪心
func jump3(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	maxPoint := 0
	targetPoint := nums[0]
	jump := 0
	for i := 0; i < len(nums); i++ {
		if i+nums[i] > maxPoint {
			maxPoint = i + nums[i]
		}
		if i == targetPoint || i == len(nums)-1 {
			jump++
			targetPoint = maxPoint
		}
	}
	return jump
}
