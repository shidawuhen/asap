/**
@author: Jason Pang
@desc:
@date: 2022/1/31
**/
package algorithm

import (
	"fmt"
	"sort"
)

/**
 * 原题：https://leetcode-cn.com/problems/3sum/submissions/
	解题思路：
	方案一：
	数组排序后，选取两个数值，剩下的数据从map中选择，如果有对应值，则判断是否已经存在，如果不存在则记录，否则不记录。
	这个方案导致超时，所以需要进一步优化。如果三个数据的和为0，则至少有一个为负值，一个为正值，所以遍历的时候，第一层遍历，只需要遍历到正值开始之前即可。
	另外，只能说Go的执行效率够高。

*/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	res := make(map[string][]int)
	sort.Ints(nums)
	positivePos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 && nums[i-1] < 0 {
			positivePos = i
		}
	}
	//fmt.Println(nums)
	vMap := make(map[int]map[int]bool) //值/位置
	for p, v := range nums {
		if _, ok := vMap[v]; ok {
			vMap[v][p] = true
		} else {
			vMap[v] = map[int]bool{
				p: true,
			}
		}
	}
	//fmt.Println(vMap)
	for i := 0; i <= positivePos; i++ {
		for j := i + 1; j < len(nums); j++ {
			v := -(nums[i] + nums[j])
			// fmt.Println(v,vMap[v])
			if m, ok := vMap[v]; ok { //存在对应值
				//判断是否为与i、j重合
				sameNum := 1
				if nums[i] == nums[j] {
					sameNum = 2
				}
				if posNotExists(i, m, sameNum) && posNotExists(j, m, sameNum) { //位置不存在，则记录
					k := ""
					var arr []int
					if v <= nums[i] {
						k = fmt.Sprintf("%d_%d_%d", v, nums[i], nums[j])
						arr = []int{v, nums[i], nums[j]}
					} else if v > nums[i] && v <= nums[j] {
						k = fmt.Sprintf("%d_%d_%d", nums[i], v, nums[j])
						arr = []int{nums[i], v, nums[j]}
					} else if v > nums[j] {
						k = fmt.Sprintf("%d_%d_%d", nums[i], nums[j], v)
						arr = []int{nums[i], nums[j], v}
					}
					//fmt.Println(i,j,k,nums[i],nums[j],v)
					if _, ok := res[k]; !ok {
						res[k] = arr
					}
				}
			}
		}
	}
	// fmt.Println(res,len(res))
	fin := make([][]int, 0)
	for _, item := range res {
		fin = append(fin, item)
	}
	return fin
}

/**
 * @Author: Jason Pang
 * @Description: 位置不存在
 * @param pos
 * @param posMap
 * @return bool
 */
func posNotExists(pos int, posMap map[int]bool, sameNum int) bool {
	if _, ok := posMap[pos]; !ok { //不存在
		return true
	}
	if len(posMap) > sameNum {
		return true
	}
	return false
}
