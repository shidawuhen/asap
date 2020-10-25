package algorithm

import "strings"

/*
原题：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."


限制：

0 <= s 的长度 <= 10000

通过次数101,808提交次数134,328
*/

func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
