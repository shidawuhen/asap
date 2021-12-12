/**
@author: Jason Pang
@desc:
@date: 2021/12/12
**/
package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/integer-to-roman/
12. 整数转罗马数字
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。



示例 1:

输入: num = 3
输出: "III"
示例 2:

输入: num = 4
输出: "IV"
示例 3:

输入: num = 9
输出: "IX"
示例 4:

输入: num = 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.
示例 5:

输入: num = 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.


提示：

1 <= num <= 3999
*/
/**
 * @Author: Jason Pang
 * @Description:
思路：
1.整理好罗马数字显示的规律
2.将整数拆分，把数字往规律上映射
 * @param num
 * @return string
*/
func intToRoman(num int) string {
	mapIToR := map[int][3]string{
		1:    [3]string{"I", "V", "X"},
		10:   [3]string{"X", "L", "C"},
		100:  [3]string{"C", "D", "M"},
		1000: [3]string{"M", "M", "M"},
	}
	//看是到哪个位置，1~3拼[0]，4拼，5拿，6~8拼，9拼
	fmt.Println(mapIToR)
	roman := ""
	multi := 1
	for num != 0 {
		s := ""
		left := num / 10
		v := num - left*10 //当前位置上的值
		symbol, _ := mapIToR[multi]
		if v <= 3 {
			for i := 0; i < v; i++ {
				s += symbol[0]
			}
		} else if v == 4 {
			s = symbol[0] + symbol[1]
		} else if v == 5 {
			s = symbol[1]
		} else if v >= 6 && v <= 8 {
			s = symbol[1]
			for i := 6; i <= v; i++ {
				s += symbol[0]
			}
		} else if v == 9 {
			s = symbol[0] + symbol[2]
		}
		roman = s + roman
		num = left
		multi *= 10
	}
	return roman
}
