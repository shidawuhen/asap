package algorithm

/*
原题：https://leetcode-cn.com/problems/valid-anagram/
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。


*/

func isAnagram(s string, t string) bool {
	record := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		record[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if _, ok := record[t[i]]; ok {
			record[t[i]]--
		} else {
			return false
		}
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}
