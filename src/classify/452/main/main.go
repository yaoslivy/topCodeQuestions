package main

import "sort"

func findMinArrowShots2(points [][]int) int {
	//先按右端点从小到大排序，保证了遍历时只需要判断当前区间的右端点是否小于下一区间的左端点
	sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
	// 从左到右依次找出最大的覆盖区间
	right := points[0][1] //记录当前覆盖区间的最大值
	res := 1
	for i := 1; i < len(points); i++ {
		if right < points[i][0] { //如果下一个区间的和当前区间不重叠
			res++                //实际上计算出了非重叠区间的个数，前一个重叠区间需要一箭引爆
			right = points[i][1] //更新当前区间的右端点
		}
	}
	return res
}

func findMinArrowShots(points [][]int) int {
	// 局部最优：当气球重叠时，一起
	// 全局最优：把所有气球射爆所用弓箭最少
	// 先对气球按照起始位置排序，如果气球重叠了，
	// 重叠气球中右边界的最小值之前的区间一定需要一个弓箭，不然靠左的重叠气球就不能被引爆
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	res := 1              //记录弓箭数量
	right := points[0][1] //记录当前重叠区间右边界
	for i := 1; i < len(points); i++ {
		if right < points[i][0] { // 当前重叠区间右边界无法覆盖下一个气球
			res++
			right = points[i][1]
		} else {
			right = min(right, points[i][1])
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
