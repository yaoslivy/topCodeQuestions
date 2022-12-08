package main

func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}
	//统计i-j是否是回文子串
	arrIsValid := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		arrIsValid[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (i+1 >= j || arrIsValid[i+1][j-1]) {
				arrIsValid[i][j] = true
			}
		}
	}
	//dp[i] 表示0到i-1位置的最小分割次数
	//dp[i] = min(dp[i], dp[j]+1) //if arrIsValid[j+1][i] //j \in 0-i
	//dp[i] = i
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = i
	}
	for i := 1; i < len(s); i++ {
		//整个字符串为回文串
		if arrIsValid[0][i] {
			dp[i] = 0
			continue
		}
		//需要切割，寻找切割点
		for j := 0; j < i; j++ {
			if arrIsValid[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(s)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
