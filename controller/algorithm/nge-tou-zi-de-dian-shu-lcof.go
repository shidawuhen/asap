package algorithm

import "math"

/*
原题：https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/
剑指 Offer 60. n个骰子的点数
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。



你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。



示例 1:

输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
示例 2:

输入: 2
输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]


限制：

1 <= n <= 11

分析：动态规划法
举个例子，2个骰子算出4点的概率g(2,4)=g(1,1)+g(1,2)+g(1,3)
*/
func dicesProbability(n int) []float64 {
	record := make([][]int, n)
	record[0] = make([]int, 7)
	for i := 1; i < 7; i++ {
		record[0][i] = 1
	}
	for i := 1; i < n; i++ { //几个骰子
		num := i + 1
		sIndex := num
		eIndex := num * 6
		record[i] = make([]int, eIndex+1)
		//fmt.Println(sIndex,eIndex)
		for j := sIndex; j <= eIndex; j++ { //几个骰子的合
			v := 0
			sm := int(math.Max(float64(num-1), float64(j-6)))
			//em := int(math.Min(float64((num - 1)*6),float64(j)))
			em := j
			if em > (num-1)*6 {
				em = (num-1)*6 + 1
			}
			//fmt.Println(j,sm,em)
			for m := sm; m < em; m++ {
				v += record[i-1][m]
			}
			record[i][j] = v
		}
		//fmt.Println(record[i])
	}
	res := record[n-1][n : n*6+1]
	sum := 0
	for i := 0; i < len(res); i++ {
		sum += res[i]
	}
	fin := make([]float64, len(res))
	for i := 0; i < len(res); i++ {
		fin[i] = float64(res[i]) / float64(sum)
	}
	//fmt.Println(record)
	return fin
}
