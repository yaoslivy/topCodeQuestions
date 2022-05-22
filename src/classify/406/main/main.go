package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	//先按身高从大到小排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	//[[7 0] [7 1] [6 1] [5 0] [5 2] [4 4]]
	// 局部最优：优先按高的people的k来插入，插入操作后已入队的元素满足条件
	// 全局最优：最后做完插入操作，整个队列满足队列属性
	res := make([][]int, 0)
	for _, val := range people {
		pos := val[1]
		res = append(res[:pos], append([][]int{val}, res[pos:]...)...)
	}
	return res
}
