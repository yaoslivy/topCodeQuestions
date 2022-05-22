package main

import "math"

func main() {

}

func numSquares(n int) int {
	//等价于用物品填满背包容量为n所用的最少物品数，且物品可以重复使用
	// dp[j] 和为j的数的组合方法数
	// dp[j] += dp[j-i*i] + 1
	// dp[j] = math.MaxInt dp[0] = 0
	dp := make([]int, n+1)
	dp[0] = 0 //
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i*i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j-i*i >= 0 && dp[j-i*i] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	if dp[n] == math.MaxInt {
		return -1
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}