package algorithm

/*
原题：https://leetcode-cn.com/problems/longest-palindromic-substring/
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) string {
	maxL := 0
	length := len(s)
	var record string
	for i := 0; i < length; i++ {
		l, r := i-1, i+1
		sl := 1
		//计算奇数个
		for l >= 0 && r < length && s[l] == s[r] {
			sl += 2
			l--
			r++
		}
		if sl > maxL {
			maxL = sl
			record = s[l+1 : r]
		}
		//计算偶数个
		sl = 0
		l, r = i, i+1
		for l >= 0 && r < length && s[l] == s[r] {
			l--
			r++
			sl += 2
		}
		if sl > maxL {
			maxL = sl
			record = s[l+1 : r]
		}
	}
	return record
}
