package algorithm

/*
原题：https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
剑指 Offer 66. 构建乘积数组
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。



示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]


提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000

分析：
画图找规律，可以发现构成对称三角形。
左下三角形，下一个比上一个多乘以一个数据
右上三角形同理
*/
func constructArr(a []int) []int {
	if len(a) == 0 {
		return a
	}
	b := make([]int, len(a))
	b[0] = 1
	//计算下三角
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	//计算上三角
	tmp := 1
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}
