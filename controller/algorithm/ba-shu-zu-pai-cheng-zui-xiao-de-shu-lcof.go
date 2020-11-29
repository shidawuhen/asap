package algorithm

import (
	"sort"
	"strconv"
)

/*
原题：https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
剑指 Offer 45. 把数组排成最小的数
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。



示例 1:

输入: [10,2]
输出: "102"
示例 2:

输入: [3,30,34,5,9]
输出: "3033459"


提示:

0 < nums.length <= 100
说明:

输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/

type MinNumberData []int

func (m MinNumberData) Len() int {
	return len(m)
}
func (m MinNumberData) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m MinNumberData) Less(i, j int) bool {
	//转换为string
	iS := strconv.Itoa(m[i])
	jS := strconv.Itoa(m[j])
	s1 := iS + jS
	s2 := jS + iS
	if s1 <= s2 {
		return true
	} else {
		return false
	}
}

func minNumber(nums []int) string {
	m := MinNumberData{}
	m = nums
	sort.Sort(m)
	//fmt.Println(m)
	s := ""
	for i := 0; i < len(m); i++ {
		s += strconv.Itoa(m[i])
	}
	return s
}
