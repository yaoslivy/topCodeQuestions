package main

import "fmt"

func main() {
	m := 19
	n := 13
	fmt.Println(uniquePaths(m, n))
}

func uniquePaths2(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	//dp[i][j]表示从0，0到i，j位置的不同路径数目
	// dp[i][j] = dp[i-1][j] + dp[i][j-1] 由上一个点往下走，或是由左边一个点往右走一步
	//初始化
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

var res int

func uniquePaths(m int, n int) int {
	if m <= 1 && n <= 1 {
		return 1
	}
	res = 0
	find(m, n, 1, 1)

	return res
}

func find(m, n int, row int, col int) {
	if row == m && col == n {
		res++
		return
	}
	//向下
	if row+1 <= m {
		find(m, n, row+1, col)
	}
	//向右
	if col+1 <= n {
		find(m, n, row, col+1)
	}

}
