package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	res := make([][]int, 0)
	// 按左边界从小到大排序
	// 局部最优：每次合并都取最大的右边界，这样就可以合并更多的区间
	// 全局最优：合并所有重叠区区间
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] { //更新右边界
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else { //不重叠直接加入
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge2(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	res := make([][]int, 0)
	// 按左边界从小到大排序
	// 局部最优：每次合并都取最大的右边界，这样就可以合并更多的区间
	// 全局最优：合并所有重叠区区间
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	flag := false //标记最后一个区间是否合并
	// 1,3  2,6
	for i := 1; i < len(intervals); i++ {
		start := intervals[i-1][0] //初始范围
		end := intervals[i-1][1]
		for i < len(intervals) && intervals[i][0] <= end { //合并区间
			end = max(end, intervals[i][1]) //不断更新右区间
			if i == len(intervals)-1 {
				flag = true
			}
			i++
		}
		res = append(res, []int{start, end})
	}
	if flag == false { //最后一区间没有合并
		res = append(res, []int{intervals[len(intervals)-1][0], intervals[len(intervals)-1][1]})
	}
	return res
}
