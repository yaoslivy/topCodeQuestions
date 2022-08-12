package main

func maxProfit(prices []int, fee int) int {
	//状态：持有股票时的金额，卖出股票时的金额
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee) //前一天卖出状态；前一天买入状态，今天卖出
	// dp[0][0] = -prices[0] dp[0][1] = 0
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[len(prices)-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit2(prices []int, fee int) int {
	// 每次在最低点买入，在最高点卖出的差值diff -fee就是一段买卖的最大利率,
	// diff-fee > 0才收集利润
	minPrice := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		// 如果当前的价格小于最低价，则在此处买入
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		// 可以收集利润的情况
		if prices[i]-minPrice-fee > 0 {
			res += prices[i] - minPrice - fee
			//更新当前的最低点，minVal到i+1是一段上升序列，收集了利润
			//如果还有下一点还可以收集利润，则不能减去fee
			minPrice = prices[i] - fee //只有当最低价格变化时，才减去fee
		}

	}
	return res
}
