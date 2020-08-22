package algorithm

/*
原题
https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
剑指 Offer 53 - II. 0～n-1中缺失的数字
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
限制：
1 <= 数组长度 <= 10000

分析：
1.方法一，长度太小，不使用二分，直接暴力查
2.这道题有个隐藏的坑，如果n=4，数组为0 1 2 3，需要显示4。但是题目描述的并不准确。
*/

func MissingNumber(nums []int) int {
	for k, v := range nums {
		if k != v {
			return k
		}
	}
	//如果执行到这，表明上面全都一致，所以去nums最后一位加1
	return nums[len(nums)-1] + 1
}

/*
方法二：二分法。还是按照相对高端一点的算法来做吧
*/

func MissingNumberDichotomize(nums []int) int {
	/*numLength := len(nums)
	leftIndex := 0
	rightIndex := numLength - 1
	for i := 0; i < numLength; i++ {
		//结束判断
		if leftIndex == rightIndex {
			if nums[leftIndex] == leftIndex {
				return leftIndex + 1
			} else { //如果不一致，则索引便是缺失的值
				return leftIndex
			}
		}
		//查找中间元素
		middleIndex := (rightIndex + leftIndex) / 2
		if nums[middleIndex] == middleIndex { //如果一致，说明前半部分没有问题
			if middleIndex+1 <= rightIndex { //确保左侧移动过程中不超过rightIndex
				leftIndex = middleIndex + 1
			} else {
				leftIndex = middleIndex
			}
		} else { //不相等表明middleIndex所指的位置已经可能是丢失的位置，保持rightIndex不+-1移动，leftIndex+-1移动
			rightIndex = middleIndex
		}
	}
	return nums[len(nums)-1] + 1
	*/
	leftIndex := 0
	rightIndex := len(nums) - 1
	midIndex := (leftIndex + rightIndex) / 2
	for leftIndex <= rightIndex {
		if nums[midIndex] == midIndex {
			leftIndex = midIndex + 1
		} else {
			rightIndex = midIndex - 1
		}
		midIndex = (leftIndex + rightIndex) / 2
	}
	return leftIndex
}
