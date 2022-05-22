package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 3, 3, 4, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	// 等价于在len(nums)个物品中找到sum/2
	// dp[j] 最接近j的物品组合和
	// dp[j] = max(dp[j], dp[j-nums[i]] + nums[i]) //1至i-1件物品中的最接近组合 使用第i件物品的最接近组合
	// dp[0] = 0
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 { //不存在均分元素的情况
		return false
	}
	aim := sum / 2
	dp := make([]int, aim+1)
	for i := 0; i < len(nums); i++ {
		for j := aim; j-nums[i] >= 0; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[aim] == aim { //刚好凑满则可以使得两个子集的元素和相等
		return true
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var isTrue bool

func canPartition2(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	sort.Ints(nums)
	if sum%2 != 0 {
		return false
	}
	isTrue = false
	find(sum/2, nums, 0, []int{})
	return isTrue
}

func find(sum int, nums []int, start int, oneRes []int) {
	if len(oneRes) != 0 {
		res := 0
		for _, val := range oneRes {
			res += val
		}
		if res == sum {
			isTrue = true
			return
		}
	}

	for i := start; i < len(nums); i++ {
		oneRes = append(oneRes, nums[i])
		find(sum, nums, i+1, oneRes)
		if isTrue == true {
			return
		}
		oneRes = oneRes[:len(oneRes)-1]
	}
}
