package algorithm

/*
原题：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func LengthOfLongestSubstring(s string) int {
	maxLength := 0
	length := len(s)
	if length == 0 {
		return maxLength
	}
	l := 0
	record := make(map[byte]int)
	for i := 0; i < length; i++ {
		if _, ok := record[s[i]]; !ok {
			l++
		} else {
			sameIndex, _ := record[s[i]]
			for key, index := range record {
				if index <= sameIndex {
					delete(record, key)
				}
			}
			l = len(record) + 1
		}
		//fmt.Println(l,maxLength,record,string(s[i]))
		record[s[i]] = i
		if l > maxLength {
			maxLength = l
		}
	}
	//fmt.Println(maxLength)
	return maxLength
}
