package main

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
