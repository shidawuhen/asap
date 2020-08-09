package algorithm
/*
原题:
https://leetcode-cn.com/problems/two-sum/
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

分析：
1.可以使用减法，反向计算数值是否存在
2.为了快速找到数值，可以使用hashmap
3.因为提到每种输入只会对应一个答案，所以意味相同的数值最多有两个。
因此建两个hashmap，一个存放所有值(重复的放另一个hashmap里），一个存放第二次重复的值
4.题目没有说大小，先不管内存问题
*/
func TwoSum(nums []int, target int) []int {
	record1 := make(map[int]int)
	recore2 := make(map[int]int)
	res := make([]int,2)
	//将nums中的值映射到hashmap中
	for k,v := range nums {
		if _,ok := record1[v];!ok{
			record1[v] = k
		}else{
			recore2[v] = k
		}
	}
	for number,index := range record1 {
		diff := target - number
		r1Index, ok1 := record1[diff]
		if ok1 {
			if r1Index != index {
				res[0] = index
				res[1] = r1Index
				break
			}else{
				r2Index, ok2 := recore2[diff]
				if ok2{
					res[0] = index
					res[1] = r2Index
					break
				}
			}
		}
	}
	if res[0] > res[1]{
		res[0],res[1] = res[1],res[0]
	}
	return res
}
