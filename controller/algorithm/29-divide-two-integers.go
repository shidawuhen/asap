/**
@author: Jason Pang
@desc:
@date: 2022/1/14
**/
/*
原题：https://leetcode-cn.com/problems/divide-two-integers/

29. 两数相除
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2



示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2


提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/
package algorithm

import (
	"fmt"
	"math"
)

//方案1：一个一个不断增加，超时
func divideMethod1(dividend int, divisor int) int {
	//divisor不断增加，当第一次大于等于被除数的时候，意味找到了数量
	count := 1
	res := divisor
	for res >= dividend && dividend-res <= divisor {
		res += divisor
		count++
	}
	return count
}

// X/Y=Z ->  Y*Z=X
//商Z乘以除数Y，商Z二进制化
//比X小，意味超过了，商太大，返回false，需要让商小一点
//快速加
//不能超，因为超了就多加了一个1
func quickAdd(x, y, z int) bool {
	zb := fmt.Sprintf("%b", z)
	//从低位到高位
	res := 0
	add := y
	for i := len(zb) - 1; i >= 0; i-- {
		if string(zb[i]) == "1" { //需要相加
			if x-res > add {
				return false
			}
			res += add
		}
		if i != 0 { //已经到头了，不需要再判断
			if x-add > add {
				return false
			}
			add += add
		}
	}
	return true
}

//方案2：二分查找
func divideMethod2(dividend int, divisor int) int {
	ans := 0
	left, right, mid := 1, math.MaxInt32, 0
	for left <= right {
		mid = left + (right-left)>>1
		if quickAdd(dividend, divisor, mid) {
			ans = mid
			if mid == math.MaxInt32 { //防止溢出
				return mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func divide(dividend int, divisor int) int {
	//只有被除数为最小值，除数为-1的时候，才会溢出，此时返回最大值
	if dividend == math.MinInt32 {
		if divisor == -1 {
			return math.MaxInt32
		}
		if divisor == 1 {
			return math.MinInt32
		}
	}

	//为简化计算，将数据改为相同符号；需要都改为负号，因为如果都改为正号，如果输入为-2^31会溢出
	reverse := false
	if dividend > 0 {
		dividend = -dividend
		reverse = !reverse
	}
	if divisor > 0 {
		divisor = -divisor
		reverse = !reverse
	}

	//被除数比除数大，则值必然比1小，返回0
	if divisor < dividend {
		return 0
	}

	//方案1：一个一个不断增加，超时
	//count := divideMethod1(dividend, divisor)
	//方案2：二分查找
	count := divideMethod2(dividend, divisor)

	if reverse {
		count = -count
	}
	return count
}
