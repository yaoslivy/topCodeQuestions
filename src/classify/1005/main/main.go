package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	// 局部最优：让绝对值大的负数变成正数，当k依然大于0时，再将绝对值小的数取反
	// 全局最优：整个数组和达到最大
	//先将整个序列按照绝对值从大到小排序，
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i]))-math.Abs(float64(nums[j])) > 0
	})

	sum := 0
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] *= -1
			k--
		}
	}
	//若序列中都为正数，且还有k未使用完，奇数则将末尾绝对值最小的数转换为负数，偶数不需要处理，因为两次转换效果相同
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}

	for _, val := range nums {
		sum += val
	}

	return sum
}
