package algorithm

import (
	"math"
	"strconv"
)

/*
原题：https://leetcode-cn.com/problems/subsets/
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

分析：
该数组的子集和同长度的比特串的子集是一一对应的
000
001
010
011
100
101
110
111
这是从0~2^n之间的数值，所以通过从0~2^n之间的数值计算比特串，为0表示该位置的数不要，1为要

有一个注意点，就是比特串的顺序需要从右往左遍历，否认在1和10的效果会是一致的
如果从右往左遍历，nums的顺序也需要在反转一次，反反得正
*/

// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ;num > 0 ; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}


func Subsets(nums []int) [][]int {
	length := len(nums)
	//计算2^n
	totalArray := int(math.Pow(2,float64(length)))
	//申请二维数组
	res := make([][]int,totalArray)
	for i := 0; i < totalArray; i++{
		//计算比特串
		bitString := convertToBin(i)
		//计算出子串中1的数量
		count1 := 0
		for j := 0; j<len(bitString);j++{
			if bitString[j] != byte('0') {
				count1++
			}
		}
		//申请子集大小
		res[i] = make([]int,count1)
		index := 0
		//反向遍历比特串，如果二级制为1，则记录，切记nums也需要反向
		for j := len(bitString) - 1; j >= 0;j--{
			if bitString[j] != byte('0') {
				res[i][index] = nums[len(bitString) - 1-j]
				index++
			}
		}
	}
	return res
}
