package main

func minDistance(word1 string, word2 string) int {
	// dp[i][j] 表示以word1[i-1]结尾的子串和以word2[j-1]结尾的子串的最近编辑距离
	// dp[i][j] = dp[i-1][j-1] if word1[i-1] == word2[j-1]
	// = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1) //word1删除一个元素，word2删除一个元素， word1替换一个元素，
	// dp[i][0] = i dp[0][j] = j dp[0][0] = 0
	dp := make([][]int, len(word1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for j := 1; j < len(word2)+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
