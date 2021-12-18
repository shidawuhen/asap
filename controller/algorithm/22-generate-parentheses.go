/**
@author: Jason Pang
@desc:
@date: 2021/12/17
**/
package algorithm

/*
原题：https://leetcode-cn.com/problems/generate-parentheses/
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]


提示：

1 <= n <= 8
*/

var allParenthesis map[int][]string

func generateParenthesis(n int) []string {
	allParenthesis = make(map[int][]string)
	return generate(n)
}

func generate(n int) []string {
	if n <= 0 {
		return []string{""}
	}
	if l, ok := allParenthesis[n]; ok {
		return l
	}
	allParenthesis[n] = make([]string, 0)
	for i := 0; i < n; i++ {
		fa := generate(i)
		fb := generate(n - i - 1)
		for _, faItem := range fa {
			for _, fbItem := range fb {
				allParenthesis[n] = append(allParenthesis[n], "("+faItem+")"+fbItem)
			}
		}
	}
	return allParenthesis[n]
}
