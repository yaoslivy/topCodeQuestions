package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口:
	// 窗口内是什么？ 满足其和>=aim的长度最小的连续子数组
	// 如何移动窗口的起始位置？  窗口的值总和>aim，就可以缩小了
	// 如何移动窗口的结束位置？  遍历数组指针，到边界停止移动
	i, j := 0, 0
	sum := 0
	res := math.MaxInt
	for j < len(nums) {
		sum += nums[j]
		for sum >= target { //调整起始位置
			res = min(res, j-i+1) //满足条件的结果值
			sum -= nums[i]
			i++
		}
		j++
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
