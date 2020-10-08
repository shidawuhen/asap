package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/task-scheduler/
621. 任务调度器
给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，
并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。

然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的最短时间。



示例 ：

输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
     在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。

提示：

任务的总个数为 [1, 10000]。
n 的取值范围为 [0, 100]。

分析：
1. 将任务按照多少排序
2. 如果不同任务量大于等于n，则直接将这些符合的减去
3. 如果不同任务量小于n，则用待命填充

其实按照最大限度填充即可，因为多余出来的按照列往上顶就可以

*/

func LeastInterval(tasks []byte, n int) int {
	sum := 0
	length := len(tasks)
	if length == 0 {
		return sum
	}
	t := make([]int, 26)
	for i := 0; i < length; i++ {
		t[tasks[i]-'A']++
	}
	quickSort(t, 0, 25)
	//fmt.Println(t)
	for i := 0; i < 13; i++ {
		t[i], t[26-1-i] = t[26-1-i], t[i]
	}
	//fmt.Println(t)
	for t[0] > 1 {
		tag := 0
		for i := 0; i < 26; i++ {
			if t[i] > 0 {
				t[i]--
				sum++
				tag++
				length--
			}
			if tag == n+1 {
				break
			}
		}
		sum += (n + 1) - tag
		quickSort(t, 0, 25)
		//fmt.Println(t)
		for i := 0; i < 13; i++ {
			t[i], t[26-1-i] = t[26-1-i], t[i]
		}

	}
	for i := 0; i < 26; i++ {
		sum += t[i]
	}

	fmt.Println(t, sum)
	return sum
}
