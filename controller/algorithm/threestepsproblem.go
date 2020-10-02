package algorithm

/*
原题：https://leetcode-cn.com/problems/three-steps-problem-lcci/submissions/
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

示例1:

 输入：n = 3
 输出：4
 说明: 有四种走法
示例2:

 输入：n = 5
 输出：13
提示:

n范围在[1, 1000000]之间

分析：
这道题比较简单，但是也算是符合动态规划策略。举个简单的例子，如果现在有5个台阶，你知道4个台阶时有多少种可能
3个台阶时多少种可能，2个台阶时多少种可能，那么从4到5只能走1阶，从3到5只能走2阶，因为走1阶的这种已经在4个台阶时包含了，
所以递推公式为:f(n)=f(n-1)+f(n-2)+f(n-3)
这道题需要注意的地方在于每计算出一个f(n)，需要记录取模的值
*/

func waysToStep(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	val := make([]int, n+1)
	val[1] = 1
	val[2] = 2
	val[3] = 4
	for i := 4; i <= n; i++ {
		val[i] = (val[i-1] + val[i-2] + val[i-3]) % 1000000007
	}
	return val[n]
}

//爬梯子 https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	val := make([]int, n+1)
	val[1] = 1
	val[2] = 2
	for i := 3; i <= n; i++ {
		val[i] = val[i-1] + val[i-2]
	}
	return val[n]
}
