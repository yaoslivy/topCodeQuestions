package main

func maxProfit(prices []int) int {
	//至多买卖两次，可能的情况：可以买卖一次，买卖两次和不买卖
	//dp[i][j] 第i天的状态j下所剩的最大现金，j: 0没有操作，1第一次买入，2第一次卖出，3第二次买入，4第二次卖出
	//dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	//dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
	//dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
	//dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	//dp[0][0] = 0; dp[0][1] = -prices[0]; dp[0][2] = 0; dp[0][3] = -prices[0]; dp[0][4] = 0
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 5)
	}
	dp[0][1], dp[0][3] = -prices[0], -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[len(prices)-1][4]

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
