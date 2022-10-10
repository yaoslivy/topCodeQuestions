package main

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] 0至i-1位置的字符串1和0至j-1位置的字符串2的最长公共子序列
	// dp[i][j] = dp[i-1][j-1] + 1 if text1[i-1] == text2[j-1]
	// else  = max(dp[i-1][j], dp[i][j-1]) //去掉i-1位置字符和去掉j-1位置字符
	// dp[i][0] = 0 dp[0][j] = 0
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
