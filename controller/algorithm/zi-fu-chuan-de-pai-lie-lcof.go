package algorithm

/*
原题：https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
剑指 Offer 38. 字符串的排列
输入一个字符串，打印出该字符串中字符的所有排列。



你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。



示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]


限制：

1 <= s 的长度 <= 8


*/
func permutation(s string) []string {
	length := len(s)
	if length == 0 {
		return []string{}
	}
	//前n-1个字符排序好后，把第n个放入每个空隙中，包含收尾
	newArrayLen := 1
	oldArray := make([]string, 1)
	oldArray[0] = string(s[0])
	for i := 1; i < length; i++ {
		newArrayLen = newArrayLen * (i + 1) //n!
		newArray := make([]string, newArrayLen)
		index := 0
		for j := 0; j < len(oldArray); j++ { //对于n-1个值的全排列，将第n个字符放到每一个间隙里
			currentS := oldArray[j]
			for k := 0; k <= len(currentS); k++ {
				newArray[index] = currentS[0:k] + string(s[i]) + currentS[k:len(currentS)]
				index++
			}
		}
		oldArray = newArray
	}
	record := make(map[string]bool)
	newArray := make([]string, 0)
	for i := 0; i < len(oldArray); i++ {
		if _, ok := record[oldArray[i]]; !ok {
			newArray = append(newArray, oldArray[i])
			record[oldArray[i]] = true
		}
	}
	return newArray
}
