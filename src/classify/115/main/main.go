package main

func numDistinct(s string, t string) int {
	// dp[i][j] 表示0至i-1序列s 和 出现以元素t[j-1]结尾的序列t的个数
	// dp[i][j] = dp[i-1][j-1] + dp[i-1][j] if s[i-1] == t[j-1] //s以最后一个元素结尾和不以最后一个元素结尾
	// else  = dp[i-1][j]
	// dp[0][0] = 1 dp[i][0] = 1 dp[0][j] = 0
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] { //相等时需要考虑是否将s[i-1]作为子串, 比如：bagg和bag
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
