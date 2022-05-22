package main

import "sort"

func findMinArrowShots(points [][]int) int {
	// 局部最优：当气球重叠时，一起
	// 全局最优：把所有气球射爆所用弓箭最少
	// 先对气球按照起始位置排序，如果气球重叠了，重叠气球中右边界的最小值之前的区间一定需要一个弓箭，不然靠左的重叠气球就不能被引爆
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] { //i气球的左边界不在i-1气球的范围内，不重叠，先射爆i-1
			res++
		} else { //重叠则更新重叠气球集合的最小右边界
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
