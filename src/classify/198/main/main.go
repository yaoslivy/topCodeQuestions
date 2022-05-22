package main

func rob(nums []int) int {
	//dp[j] 1-j间房，最多可以偷窃的金额
	//dp[j] = max(dp[j-2]+nums[j], dp[j-1])  是否取j间房nums[j]，取决于前一间房和前两间房的大小
	//dp[0] = nums[0] dp[1] = max(nums[0], nums[1])
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for j := 2; j < len(nums); j++ {
		dp[j] = max(dp[j-2]+nums[j], dp[j-1])
	}
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
