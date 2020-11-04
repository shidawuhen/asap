package algorithm

import (
	"math"
)

/*
原题：https://leetcode-cn.com/problems/zigzag-conversion/
6. Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING",len=16 numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/

func Convert(s string, numRows int) string {
	//第 1 4 7   0 3 6 i%(numRows-1)=0 顺着往下 否则 减一
	if numRows == 1 {
		return s
	}
	sRecord := make([][]byte, numRows)
	roundL := numRows + numRows - 2
	roundNum := int(math.Ceil(float64(len(s)) / float64(roundL)))
	for i := 0; i < numRows; i++ {
		sRecord[i] = make([]byte, roundNum*(numRows-1))
	}
	for i := 0; i*roundL < len(s); i++ {
		endIndex := (i + 1) * roundL
		if endIndex > len(s) {
			endIndex = len(s)
		}
		sUse := s[i*roundL : endIndex]
		colIndex := (numRows - 1) * i
		rowIndex := 0
		for j := 0; j < len(sUse); j++ {
			if j < numRows {
				sRecord[rowIndex][colIndex] = sUse[j]
				if j != numRows-1 {
					rowIndex++
				}
			} else {
				rowIndex--
				colIndex++
				sRecord[rowIndex][colIndex] = sUse[j]
			}
		}
	}
	res := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(sRecord[i]); j++ {
			if sRecord[i][j] != 0 {
				res += string(sRecord[i][j])
			}

		}
	}
	return res
}
