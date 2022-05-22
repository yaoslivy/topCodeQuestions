package main

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	// dp[i] 到第i个台阶需要的最低花费
	// dp[i] = min(dp[i-1], dp[i-2]) + cost[i] // 到第i个台阶一定要花费cost[i]
	// dp[0]=cost[0]
	// dp[1]=cost[1]
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
