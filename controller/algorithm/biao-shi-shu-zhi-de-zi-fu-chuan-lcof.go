package algorithm

import (
	"strconv"
	"strings"
)

/*
原题：https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
剑指 Offer 20. 表示数值的字符串
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"、"1.2E5都表示数值，
但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。

分析：
格式：A.BEC
1. 先查找是否有分割点，如.或者e，.需要在E的前面，同时.和e分别存在个数只能为1个
2. 根据e将字符串拆分为两份
*/

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	s = strings.ToLower(s)
	hasdot := 0
	hase := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			if hase > 0 {
				return false
			}
			hasdot++
		} else if s[i] == 'e' {
			hase++
		}
	}
	if hase > 1 || hasdot > 1 {
		return false
	}
	arrS := strings.Split(s, "e")
	for i := 0; i < len(arrS); i++ {
		_, e := strconv.ParseFloat(arrS[i], 64)
		if e != nil {
			return false
		}
	}
	return true
}
