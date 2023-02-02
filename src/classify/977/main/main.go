package main

import (
	"math"
	"sort"
)

func sortedSquares(nums []int) []int {
	//递增序列，平方后最大值一定在两边去的，设置双指针观察平方后最大值取值
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	k := len(nums) - 1
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] { //头指针指向的值平方后更大
			res[k] = nums[i] * nums[i]
			k--
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
			k--
		}
	}
	return res
}

func sortedSquares2(nums []int) []int {
	// 按照绝对值从小到大排序
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	})
	// 将每个位置的值平方
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	return nums
}
