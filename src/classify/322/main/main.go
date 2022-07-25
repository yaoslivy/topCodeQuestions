package main

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	//dp[j] 凑成金额j所使用的最少硬币个数
	//dp[j] = min(dp[j], dp[j-coins[i]]+1)
	//0至i-1凑成的最少硬币个数 使用第i个硬币凑成的最少硬币个数=dp[j-coins[i]] + 1
	//dp[j] = math.MaxInt dp[0] = 0
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 && dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] != math.MaxInt {
		return dp[amount]
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
