package main

func longestPalindromeSubseq(s string) int {
	// dp[i][j] 表示在[i, j]范围内最长的回文子序列的长度
	// dp[i][j] = dp[i+1][j-1]+2 if s[i] == s[j]
	// else max(dp[i+1][j], dp[i][j-1]) // 不相等则观察加入哪一个的长度更大
	// dp[i][i] = 1
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
