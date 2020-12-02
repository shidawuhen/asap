package algorithm

/*
原题：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
剑指 Offer 50. 第一个只出现一次的字符
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "


限制：

0 <= s 的长度 <= 50000
*/
type Pos struct {
	index int
	num   int
}

func firstUniqChar(s string) byte {
	record := make(map[byte]Pos)
	index := 0
	for i := 0; i < len(s); i++ {
		if v, ok := record[s[i]]; !ok {
			record[s[i]] = Pos{index: index, num: 1}
			index++
		} else {
			record[s[i]] = Pos{index: v.index, num: v.num + 1}
		}
	}
	minIndex := len(s)
	var minV byte = ' '
	for k, v := range record {
		if v.num == 1 && v.index < minIndex {
			minIndex = v.index
			minV = k
		}
	}

	return minV
}
