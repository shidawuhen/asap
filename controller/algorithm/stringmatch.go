package algorithm

import (
	"math"
)

/*
原题
https://leetcode-cn.com/problems/repeated-string-match/
给定两个字符串 A 和 B, 寻找重复叠加字符串A的最小次数，使得字符串B成为叠加后的字符串A的子串，如果不存在则返回 -1。

举个例子，A = "abcd"，B = "cdabcdab"。

答案为 3， 因为 A 重复叠加三遍后为 “abcdabcdabcd”，此时 B 是其子串；A 重复叠加两遍后为"abcdabcd"，B 并不是其子串。

注意:

 A 与 B 字符串的长度在1和10000区间范围内。


分析：
1.这个主要确定终止位置
2.当多个A的长度刚刚大于等于B的长度时，最多再增加一个A就是最长的长度，因为这种情况下，继续增加只是重复
3.本题比较简单，但是字符串判断会使用KMP算法写
*/

func RepeatedStringMatch(A string, B string) int {
	s := A
	t := B
	//判断s和t的长度，将s长度扩展为N*S(刚大于t)+S
	multi := int(math.Ceil(float64(len(t))/float64(len(s)))) + 1
	finS := ""
	for i := 0; i < multi; i++ {
		finS = finS + s
	}
	startIndex := KMPCompare(finS, t)
	//fmt.Println(startIndex)
	if startIndex == -1 {
		return -1
	}
	//获得的索引+t的长度，与s长度相除，向上取整
	return int(math.Ceil(float64(startIndex+len(t)) / float64(len(s))))
}

//返回匹配开始位置的索引值，如果不匹配，返回-1
func KMPCompare(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return -1
	}
	next := computeNext(t)
	//fmt.Println(s,t,next)
	//字符串比较，以S为主进行遍历
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		//fmt.Println("i",i,"j",j,s[i],t[j])
		if s[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
			if j == -1 {
				j = 0
				i++
			}
		}
		if j == len(t) {
			return i - j
		}
	}
	return -1
}

//next[j+1]的值，根据next[j]计算出来
//next[0]=-1
func computeNext(t string) (next []int) {
	if len(t) == 0 {
		return
	}
	next = make([]int, len(t))
	next[0] = -1 //-1代表终止
	i := 1
	k := next[i-1]
	for i < len(t) {
		//如果为-1，则表明到第一个字符，肯定是最大真前缀和最大真后缀
		//如果相等，也可以直接加一
		if k == -1 || t[i-1] == t[k] {
			k++
			next[i] = k
			i++ //进入下一个字符
		} else { //k循环过程中必然最终为-1
			k = next[k]
		}
	}
	return
}
