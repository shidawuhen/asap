package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/
剑指 Offer 19. 正则表达式匹配
请实现一个函数用来匹配包含'. '和'*'的正则表达式。
模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。
注意：本题与主站 10 题相同：https://leetcode-cn.com/problems/regular-expression-matching/

"aaa"
"ab*a*c*a"
true
*/

func IsMatch(s string, p string) bool {
	pIndex := len(p) - 1
	//start := false
	fmt.Println(s, p)
	return computeZhengze(s, p, len(s)-1, pIndex)

	/*for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(i, string(s[i]), pIndex, string(p[pIndex]))
		//对于i位置，p的最后部分是否匹配
		if pIndex >= 0 {
			if p[pIndex] == '*' {
				if pIndex == 0 {
					return false
				} else {
					if p[pIndex-1] == s[i] || p[pIndex-1] == '.' {
						//start = true
					} else if p[pIndex-1] != s[i] {
						pIndex -= 2
						i++
					}
				}
			} else if p[pIndex] == s[i] || p[pIndex] == '.' {
				//start = true
				pIndex--
			} else if p[pIndex] != s[i] {
				return false
			}
		} else {
			return false
		}
	}
	fmt.Println(pIndex, p[0:pIndex+1])
	//判断p是否不存在了
	if pIndex < 0 {
		return true
	}
	for i := pIndex; i >= 0; i = i - 2 {
		if p[i] != '*' {
			return false
		} else { //p[i]==*
			if i-1 >= 0 && p[i-1] == '*' {
				i++
			}
		}
	}
	return true*/
}

func computeZhengze(s string, p string, sIndex int, pIndex int) bool {
	//fmt.Println(sIndex, pIndex)
	if pIndex < 0 && sIndex < 0 {
		return true
	}
	if pIndex < 0 && sIndex >= 0 {
		return false
	}
	if sIndex < 0 {
		for i := pIndex; i >= 0; i = i - 2 {
			if p[i] != '*' {
				return false
			} else { //p[i]==*
				if i-1 >= 0 && p[i-1] == '*' {
					i++
				}
			}
		}
		return true
	}
	//fmt.Println(sIndex, string(s[sIndex]), pIndex, string(p[pIndex]))

	//对于i位置，p的最后部分是否匹配
	if pIndex >= 0 {
		if p[pIndex] == '*' {
			if pIndex == 0 {
				return false
			} else {
				if p[pIndex-1] == s[sIndex] || p[pIndex-1] == '.' {
					//start = true
					r1 := computeZhengze(s, p, sIndex-1, pIndex)
					if r1 == true {
						return true
					}
					r2 := computeZhengze(s, p, sIndex-1, pIndex-2)
					if r2 == true {
						return true
					}
					r3 := computeZhengze(s, p, sIndex, pIndex-2)
					if r3 == true {
						return true
					}
					return false
				} else if p[pIndex-1] != s[sIndex] {
					return computeZhengze(s, p, sIndex, pIndex-2)
				}
			}
		} else if p[pIndex] == s[sIndex] || p[pIndex] == '.' {
			return computeZhengze(s, p, sIndex-1, pIndex-1)
		} else if p[pIndex] != s[sIndex] {
			return false
		}
	} else {
		return false
	}
	return false
}
