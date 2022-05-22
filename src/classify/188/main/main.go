package main

func maxProfit(k int, prices []int) int {
	//dp[i][2*k+1] 0表示不做操作的金额，1表示第一次买入，2表示第一次卖出 ...
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	// dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
	//...
	// dp[0][奇数] = -prices[0]
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 1; i < 2*k+1; i = i + 2 {
		dp[0][i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k-1; j = j + 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
