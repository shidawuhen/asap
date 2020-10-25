package algorithm

/*
原题：https://leetcode-cn.com/problems/longest-mountain-in-array/
845. 数组中的最长山脉
我们把数组 A 中符合下列属性的任意连续子数组 B 称为 “山脉”：

B.length >= 3
存在 0 < i < B.length - 1 使得 B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
（注意：B 可以是 A 的任意子数组，包括整个数组 A。）

给出一个整数数组 A，返回最长 “山脉” 的长度。

如果不含有 “山脉” 则返回 0。



示例 1：

输入：[2,1,4,7,3,2,5]
输出：5
解释：最长的 “山脉” 是 [1,4,7,3,2]，长度为 5。
示例 2：

输入：[2,2,2]
输出：0
解释：不含 “山脉”。


提示：

0 <= A.length <= 10000
0 <= A[i] <= 10000
*/

func LongestMountain(A []int) int {
	maxL := 0
	sum := 0
	length := len(A)
	if length == 0 {
		return maxL
	}
	inc := false
	dec := false
	for i := 0; i < length-1; i++ {
		if A[i] < A[i+1] { //说明是增加的
			if dec == true {
				dec = false
				if sum > maxL {
					maxL = sum
				}
				sum = 0
			}
			inc = true
			sum++
		} else if A[i] > A[i+1] { //说明是递减的
			if inc == true {
				dec = true
				sum++
				if sum > maxL {
					maxL = sum
				}
			} else {
				sum = 0
			}
		} else {
			inc = false
			dec = false
			sum = 0
		}
		//fmt.Println(A[i],A[i+1],sum,maxL,inc,dec)
	}
	//fmt.Println(maxL + 1)
	if maxL == 0 {
		return 0
	} else {
		return maxL + 1
	}
}
