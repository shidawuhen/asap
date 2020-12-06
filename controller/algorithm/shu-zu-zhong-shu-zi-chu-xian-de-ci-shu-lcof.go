package algorithm

/*
原题：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
剑指 Offer 56 - I. 数组中数字出现的次数
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。



示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
示例 2：

输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]


限制：

2 <= nums.length <= 10000

分析：
自己没想出来，需要使用异或，然后对数据进行分组。
异或的功能是一个为0一个为1，异或值为1，否则异或值为0
相同值异或结果为0，只有两个只出现一次的，至少有一个位置异或值为1.根据这个位置可以把数据分为两组
*/
func singleNumbers(nums []int) []int {
	res := 0
	//遍历查找出异或后，两个只出现一次的数字的异或结果
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	//fmt.Println(res)
	//找出第一个为1的位置，根据该位置，对数据进行分组
	pos := 1
	for (pos & res) == 0 {
		pos <<= 1
	}
	// fmt.Println(pos)
	a, b := 0, 0
	//确定分组，然后异或
	for i := 0; i < len(nums); i++ {
		//fmt.Println(nums[i],pos & nums[i])
		if (pos & nums[i]) != 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}
