package algorithm

/*
原题：https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
剑指 Offer 63. 股票的最大利润

假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。


限制：

0 <= 数组长度 <= 10^5



注意：本题与主站 121 题相同：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

分析：
v(i)表明第i天的最大利润
1）买
2）不操作
如果花费金额更小，就买
3）卖 能否赚更多

遍历数组，如果花费金额能获得股票，就在当天买
如果卖出能收益更大，就卖出


*/

func MaxProfit(prices []int) int {
	maxProfit := 0
	lenght := len(prices)
	if lenght == 0 {
		return maxProfit
	}
	minVal := prices[0]
	for i := 1; i < lenght; i++ {
		if prices[i]-minVal > maxProfit { //如果第i天卖掉能获得的最大收益
			maxProfit = prices[i] - minVal
		}
		if prices[i] < minVal { //如果第i天购买的成本更小，就在i天购买
			minVal = prices[i]
		}
	}
	//fmt.Println(maxProfit)
	return maxProfit
}
