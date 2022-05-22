package main

func maxProfit(prices []int) int {
	//dp[i][4] 0 0至i天处于买入状态，1 卖出状态，昨天为冷冻期，或昨天之前为冷冻期
	// 2 今天卖出， 3 卖出状态，今天处于冷冻期状态
	// dp[i][0] = max(dp[i-1][0], max(dp[i-1][3], dp[i-1][1])-prices[i]) //昨天买入状态，今天买入状态
	// dp[i][1] = max(dp[i-1][1], dp[i-1][3]) //昨天为冷冻期，后昨天之前为冷冻期
	// dp[i][2] = dp[i-1][0]+prices[i] // 今天卖出，昨天必须处于买入状态
	// dp[i][3] = dp[i-1][2] //昨天卖出，今天处于冷冻期
	// dp[0][0] = -prices[0]
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3], dp[i-1][1])-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[len(prices)-1][1], max(dp[len(prices)-1][2], dp[len(prices)-1][3]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
