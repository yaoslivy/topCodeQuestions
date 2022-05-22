package main

func minDistance(word1 string, word2 string) int {
	// dp[i][j] 表示0至i-1位置的子串1 和 0至j-1位置的子串2，想要到达相等时所需要删除元素的最少步数
	// dp[i][j] = dp[i-1][j-1] if word1[i-1] == word[j-1] //不需要删除
	// else = dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2 min //删除左边，右边，还是两边
	// dp[0][j] = j dp[i][0] = i
	dp := make([][]int, len(word1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+2)
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
