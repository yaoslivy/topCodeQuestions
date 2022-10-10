package main

import "math"

func maxProfit2(prices []int) int {
	//dp[i][0] 表示第1-i天持有股票所得最多现金，dp[i][1]表示第i天不持有股票所得最多现金
	//dp[i][0] = max(dp[i-1][0], -prices[i]) //第i-1天就持有股票，那就保持现状，第i天买入，则所得现金就是-prices[i]
	//dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0]) //第i-1天不持有股票，第i天不持有股票
	//dp[0][0] = -prices[0]; dp[0][1] = 0
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	}
	return dp[len(prices)-1][1]
}

func maxProfit(prices []int) int {
	lowPrice := math.MaxInt
	res := 0
	for i := 0; i < len(prices); i++ {
		lowPrice = min(lowPrice, prices[i]) //记录当前遍历序列的最小值
		res = max(res, prices[i]-lowPrice)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
