package main

func maxProfit(prices []int) int {
	// 可能状态：买入状态（之前买入，今天买入）；卖出状态（之前卖出，今天卖出）
	// dp[i][2] 表示0到i-1天两个状态下所得的最大金额
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) // 今天买入时的所得金额为前一天卖出状态下的金额+今天股票的价格
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) // 今天卖出时，前一天必须持有股票
	// dp[0][0] = -prices[0], dp[0][1] = 0
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func maxProfit2(prices []int) int {
	// 7,1,5,3,6,4
	//  -6,4,-2,3,-2
	// 在最低点买入，在最高点卖出就可以得到最大利润，这一次买卖的最大利率可以分解为相邻差值的正值集合
	// 若i为最低点，i+2为最高点， i+1为中间点，这一段买入，卖出就为获得的最大利润
	// res = prices[i+2]-prices[i] = prices[i+2] - prices[i+1] + (prices[i+1] - prices[i])
	// 求和所有正值即求和了所有买卖序列
	res := 0
	for i := 1; i < len(prices); i++ {
		dif := prices[i] - prices[i-1]
		if dif > 0 {
			res += dif
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
