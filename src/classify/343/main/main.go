package main

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	// dp[i] 表示i拆分后可以获得的最大乘积值
	// dp[i] = max(dp[i-j]*j, (i-j)*j) //拆成两部分，或者拆除三部分以上
	// dp[2] = 1
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; i-j >= 2; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
