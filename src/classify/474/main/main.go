package main

func findMaxForm(strs []string, m int, n int) int {
	//dp[i][j] 最多有i个0和j个1的strs的最大子集的大小为dp[i][j]
	//dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
	//dp[i][j] = 0
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, val := range strs { //遍历词
		zeroNum, oneNum := 0, 0
		for _, c := range val {
			if c == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		for i := m; i-zeroNum >= 0; i-- { //从后往前遍历
			for j := n; j-oneNum >= 0; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
