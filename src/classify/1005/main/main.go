package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	// 先将nums按从小到大排序，优先将负数变为正数
	// 若k次数还有剩余，剩余k为奇数才需要考虑将当前最小值变为负数
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			res -= nums[i]
			k--
		} else {
			res += nums[i]
		}
	}
	if k != 0 && k%2 == 1 {
		minVal := find(nums)
		res -= 2 * minVal
	}
	return res
}

//找到绝对值最小的值并返回
func find(nums []int) int {
	minVal := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			minVal = min(minVal, -nums[i])
		} else {
			minVal = min(minVal, nums[i])
		}
	}
	return minVal
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func largestSumAfterKNegations2(nums []int, k int) int {
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
