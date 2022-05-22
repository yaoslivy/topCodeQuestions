package main

func climbStairs(n int) int {
	// 相当于选择1，2排列和方案和n的方法数目，可以重复选择 nums[0]=1, nums[1]=2
	// dp[j] 到第j个台阶的方法数
	// dp[j] += dp[j-nums[i]]
	// dp[1]=1
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for j := 1; j <= n; j++ {
		for i := 1; i <= 2; i++ {
			if j-i >= 0 {
				dp[j] += dp[j-i]
			}
		}
	}
	return dp[n]
}

func climbStairs2(n int) int {
	dp := make([]int, 3) //dp[i]表示到达第i个台阶的方法数目

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		// dp[i] = dp[i-1] + dp[i-2]
		sum := dp[1] + dp[2]
		dp[1] = dp[2]
		dp[2] = sum
	}
	return dp[2]
}
