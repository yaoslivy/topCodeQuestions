package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//tickets := [][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}}
	tickets := [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}
	ans := findItinerary(tickets)
	fmt.Println(ans)
}

var res []string
var flag bool

func findItinerary(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		return strings.Compare(tickets[i][1], tickets[j][1]) == -1
	}) //先将目的地排序
	res = make([]string, 0)
	flag = false
	visited := make([]bool, len(tickets))
	//回溯找到第一方案就是结果
	find(tickets, []string{"JFK"}, visited)
	return res
}

func find(tickets [][]string, oneRes []string, visited []bool) {
	if len(tickets)+1 == len(oneRes) {
		res = oneRes
		flag = true
		return
	}

	for i := 0; i < len(tickets); i++ {
		//每张票在一次全排列中只使用一次；以oneRes中的终点作为起点
		if visited[i] == false && oneRes[len(oneRes)-1] == tickets[i][0] {
			oneRes = append(oneRes, tickets[i][1])
			visited[i] = true
			find(tickets, oneRes, visited)
			if flag == true { //找到的第一个即为最优
				break
			}
			visited[i] = false
			oneRes = oneRes[:len(oneRes)-1]
		}
	}
}
