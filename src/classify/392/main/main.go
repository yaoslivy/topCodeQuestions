package main

func isSubsequence(s string, t string) bool {
	// dp[i][j] 表示0至i-1序列s 和 0至j-1序列t 的的相同子序列长度
	// dp[i][j] = dp[i-1][j-1]+1 if s[i-1] == t[j-1] else dp[i][j-1] //不相同则在t中跳过，相当于比较t[j-2]之前的字符
	// dp[i][0] = 0 dp[0][j] = 0
	dp := make([][]int, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	if dp[len(s)][len(t)] == len(s) {
		return true
	}
	return false
}

func isSubsequence2(s string, t string) bool {
	// dp[i][j] 表示0到i-1序列s是否存在于0到j-1的序列t中
	// dp[i][j] = true if s[i-1] == t[j-1]  && dp[i-1][j-1] == true
	// else = dp[i][j-1]
	// dp[0][0] = true dp[0][j] = true dp[i][0] = false
	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(t)+1)
	}
	for j := 0; j < len(t)+1; j++ {
		dp[0][j] = true
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] && dp[i-1][j-1] == true {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)]
}

// 暴力匹配
func isSubsequenceTraverse(s, t string) bool {
	if len(s) == 0 {
		return true
	}
	//所有可能的起点
	for i := 0; i < len(t); i++ {
		start := i
		j := 0
		for j < len(s) && start < len(t) {
			if t[start] == s[j] {
				j++
				start++
			} else {
				start++
			}
		}
		if j == len(s) {
			return true
		}
	}
	return false
}
