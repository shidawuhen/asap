package algorithm

/*
原题：https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
剑指 Offer 51. 数组中的逆序对
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。



示例 1:

输入: [7,5,6,4]
输出: 5


限制：

0 <= 数组长度 <= 50000

分析：使用二路归并排序
这道题设计的还是很巧妙的，计算对数的方法需要仔细思考才能确定出原因
*/
func reversePairs(nums []int) int {
	cnt := reversePairsMerge(nums, 0, len(nums)-1)
	return cnt
}

func reversePairsMerge(nums []int, sIndex int, eIndex int) int {
	if sIndex >= eIndex {
		return 0
	}
	mIndex := (sIndex + eIndex) / 2

	cnt1 := reversePairsMerge(nums, sIndex, mIndex)
	cnt2 := reversePairsMerge(nums, mIndex+1, eIndex)
	tmp := make([]int, 0)
	i, j := sIndex, mIndex+1
	cnt := 0
	for i <= mIndex && j <= eIndex {
		if nums[i] <= nums[j] {
			cnt += j - (mIndex + 1)
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mIndex {
		cnt += eIndex - mIndex
		tmp = append(tmp, nums[i])
		i++
	}
	for j <= eIndex {
		tmp = append(tmp, nums[j])
		j++
	}
	//重置
	for i := sIndex; i <= eIndex; i++ {
		nums[i] = tmp[i-sIndex]
	}
	return cnt + cnt1 + cnt2
}
