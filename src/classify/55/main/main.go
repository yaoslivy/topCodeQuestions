package main

func canJump2(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	// 从左到右，不断判断所能到的最大右边界
	maxRight := nums[0]
	for i := 1; i <= maxRight; i++ {
		maxRight = max(maxRight, i+nums[i])
		if maxRight >= len(nums)-1 {
			return true
		}
	}
	return false
}

func canJump(nums []int) bool {
	// 局部最优：每次取最大跳跃步数，取最大覆盖范围
	// 全局最优：最后得到整体最大覆盖范围
	right := 0
	if len(nums) == 1 {
		return true
	}
	for i := 0; i <= right; i++ {
		right = max(i+nums[i], right) //不断更新所能覆盖的最大范围的右边界
		if right >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
