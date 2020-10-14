package algorithm

import "fmt"

/*
原题：https://leetcode-cn.com/problems/reconstruct-itinerary/

给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。
所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。

提示：

如果存在多种有效的行程，请你按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
所有的机场都用三个大写字母表示（机场代码）。
假定所有机票至少存在一种合理的行程。
所有的机票必须都用一次 且 只能用一次。


示例 1：

输入：[["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
输出：["JFK", "MUC", "LHR", "SFO", "SJC"]
示例 2：

输入：[["JFK","SFO"]3,["JFK","ATL"]1,["SFO","ATL"]4,["ATL","JFK"]2,["ATL","SFO"]5]
输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。
*/

//方法1：超时
var routes []string
var largerS string

func FindItinerary(tickets [][]string) []string {
	routes = make([]string, 0)
	largerS = ""
	length := len(tickets)
	if length == 0 {
		return routes
	}
	used := make([]int, length)
	for i := 0; i < length; i++ {
		if tickets[i][0] == "JFK" {
			used[i] = 1
			makeRoute(tickets, i, used, 1, tickets[i][0])
			used[i] = 0
		}
	}
	fmt.Println(routes)
	return routes
}

func makeRoute(tickets [][]string, currentIndex int, used []int, useNum int, routerStr string) {
	if useNum == len(tickets) {
		res := make([]string, len(used)+1)
		str := routerStr + tickets[currentIndex][1]
		for i := 0; i < len(used); i++ {
			res[used[i]-1] = tickets[i][0]
			if used[i] == len(used) {
				res[used[i]] = tickets[i][1]
			}
		}

		//fmt.Println(str)
		if str < largerS || largerS == "" {
			routes = res
			largerS = str
		}

		//fmt.Println(used,res)
	}
	for i := 0; i < len(tickets); i++ {
		if used[i] == 0 && tickets[i][0] == tickets[currentIndex][1] && (routerStr+tickets[i][0] < largerS || largerS == "") {
			used[i] = useNum + 1
			makeRoute(tickets, i, used, useNum+1, routerStr+tickets[i][0])
			used[i] = 0
		}
	}
}


//方法2：可以
/*
首先先把图的邻接表存进字典，并且按字典序排序，然后从‘JFK’开始深搜，
每前进一层就减去一条路径，直到某个起点不存在路径的时候就会跳出while循环进行回溯，
相对先找不到路径的一定是放在相对后面，所以当前搜索的起点from会插在当前输出路径的第一个位置。


var d map[string][]string
var ans []string

func findItinerary(tickets [][]string) []string {
    d = map[string][]string{}
    for _, v := range tickets {
        d[v[0]] = append(d[v[0]], v[1])
    }
    for _, v := range d {
        sort.Strings(v)
    }
    ans = []string{}
    dfs("JFK")
    n := len(ans)
	for i := 0; i < n / 2; i++ {
        ans[i], ans[n - i - 1] = ans[n - i - 1], ans[i]
	}
    return ans
}

func dfs(f string) {
    for len(d[f]) > 0 {
        v := d[f][0]
        d[f] = d[f][1: ]
        dfs(v)
    }
    ans = append(ans, f)
}

*/