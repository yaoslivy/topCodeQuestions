package main

import "math"

func maxSubArray(nums []int) int {
	// dp[i] 包括nums[i]的最大连续子序列和
	// dp[i] = max(nums[i], dp[i-1]+nums[i])
	// dp[0] = nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxSubArray3(nums []int) int {
	// 局部最优：当前连续和为负数时，立刻从下一个元素开始重新计算连续和，因为一段负数序列加上下一个值只会越来越小
	// 全局最优：在所有局部最优的情况下，找到最大值
	res := math.MinInt
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res {
			res = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return res
}

func maxSubArray2(nums []int) int {
	//暴力解法：第一层for设置起始位置，第二次for遍历数组寻找最大值
	res := math.MinInt
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			res = max(res, count)
		}
	}
	return res
}
