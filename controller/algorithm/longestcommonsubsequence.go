package algorithm

/*
原题：https://leetcode-cn.com/problems/longest-common-subsequence/
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。



示例 1:

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，它的长度为 3。
示例 2:

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。
示例 3:

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。


提示:

1 <= text1.length <= 1000
1 <= text2.length <= 1000
输入的字符串只含有小写英文字符。

分析
1. 使用动态规划
2. 使用二维数组记录结果，row为text2的每一个字符，col为text1的每一个字符
3. 如果xi=yj，则L[i][j]=L[i-1][j-1]+1
   如果xi!=yj,则L[i][j] = max{L[i-1][j],L[i][j-1]}

*/

func LongestCommonSubsequence(text1 string, text2 string) int {
	col := len(text1)
	row := len(text2)
	if col == 0 || row == 0 {
		return 0
	}
	L := make([][]int, row)
	for i := 0; i < row; i++ {
		L[i] = make([]int, col)
	}
	for i := 0; i < row; i++ { //text2
		for j := 0; j < col; j++ { //text1
			if text1[j] == text2[i] {
				lu := 0
				if i-1 >= 0 && j-1 >= 0 {
					lu = L[i-1][j-1]
				}
				L[i][j] = lu + 1
			} else {
				maxD := 0
				if i-1 >= 0 && L[i-1][j] > maxD {
					maxD = L[i-1][j]
				}
				if j-1 >= 0 && L[i][j-1] > maxD {
					maxD = L[i][j-1]
				}
				L[i][j] = maxD
			}
		}
	}
	//fmt.Println(L)
	return L[row-1][col-1]
}

/*"abcba"
 a b c b a
[1 1 1 1 1] a
[1 2 2 2 2] b
[1 2 3 3 3] c
[1 3 3 4 4] b
[1 3 4 4 4] c
[1 4 4 5 5] b
[2 4 4 5 6] a*/

/*"pmjghexybyrgzczy"  hbgc
 p m j g h e x y b y r g z c z y
[0 0 0 0 1 1 1 1 1 1 1 1 1 1 1 1]h -
[0 0 0 0 1 1 1 1 1 1 1 1 1 1 1 1]a
[0 0 0 0 1 1 1 1 1 1 1 1 1 1 1 1]f
[0 0 0 0 1 1 1 1 1 1 1 1 1 2 2 2]c -
[0 0 0 0 1 1 1 1 1 1 1 1 1 2 2 2]d
[0 0 0 0 1 1 1 1 1 1 1 1 1 2 2 2]q
[0 0 0 0 1 1 1 1 2 2 2 2 2 2 2 2]b -
[0 0 0 1 1 1 1 1 2 2 2 3 3 3 3 3]g -
[0 0 0 1 1 1 1 1 2 2 2 3 3 3 3 3]n
[0 0 0 1 1 1 1 1 2 2 2 3 3 4 4 4]c -
[0 0 0 1 1 1 1 1 2 2 3 3 3 4 4 4]r -
[0 0 0 1 1 1 1 1 2 2 3 3 3 5 5 5]c
[0 0 0 1 1 1 1 1 3 3 3 3 3 5 5 5]b
[0 0 0 1 1 1 1 1 3 3 3 3 3 5 5 5]i
[0 0 0 1 2 2 2 2 3 3 3 3 3 5 5 5]h
[0 0 0 1 2 2 2 2 3 3 3 3 3 5 5 5]k
[0 0 0 1 2 2 2 2 3 3 3 3 3 5 5 5]d*/

//"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
//"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
