package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	//先按从小到大排序， 右边界越小，留给下一个区间的空间就越大，从左向右遍历，可以优先选右边界小的
	// 局部最优： 优先选择右边界小的区间
	// 全局最优：选取最多的非交叉区间
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	// 从第二个元素开始，依次判断接下来的元素区间是否在左边的元素区间外，在这之外则更新右边界，并增加非交叉区间的个数
	right := intervals[0][1] //右边界
	res := 1                 //非交叉区间的个数
	for i := 1; i < len(intervals); i++ {
		if right <= intervals[i][0] {
			res++
			right = intervals[i][1]
		}
	}
	return len(intervals) - res
}
