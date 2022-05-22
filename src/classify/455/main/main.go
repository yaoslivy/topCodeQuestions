package main

import "sort"

func findContentChildren(g []int, s []int) int {
	//局部最优：使用小饼干喂饱小胃口的
	//全局最优：喂饱尽可能多的小孩
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
			j++
			res++
		} else {
			j++
		}
	}
	return res
}
