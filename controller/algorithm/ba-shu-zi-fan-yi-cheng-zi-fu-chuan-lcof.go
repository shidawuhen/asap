package algorithm

import (
	"fmt"
	"strconv"
)

/*
原题：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
剑指 Offer 46. 把数字翻译成字符串
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。



示例 1:

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"


提示：

0 <= num < 2^31

分析：
深度优先遍历
*/

var stringRecord []string

func translateNum(num int) int {
	stringRecord = make([]string, 0)
	var words [26]string
	for i := 0; i < 26; i++ {
		words[i] = string('a' + i)
	}
	nS := strconv.Itoa(num)
	fmt.Println(words)
	dspNum(nS, 0, words, "", len(nS))
	//fmt.Println(stringRecord)
	recordMap := make(map[string]int)
	for _, v := range stringRecord {
		if _, ok := recordMap[v]; !ok {
			recordMap[v] = 1
		}
	}
	return len(recordMap)
}

func dspNum(nS string, index int, words [26]string, s string, l int) {
	if index >= l { //最后
		stringRecord = append(stringRecord, s)
		return
	}

	v, _ := strconv.Atoi(string(nS[index]))
	dspNum(nS, index+1, words, s+words[v], l)

	if index+2 <= l && nS[index] != '0' {
		v, _ := strconv.Atoi(nS[index : index+2])
		if v < 26 {
			dspNum(nS, index+2, words, s+words[v], l)
		}
	}
}
