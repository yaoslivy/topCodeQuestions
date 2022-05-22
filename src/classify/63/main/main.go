package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	dp := make([][]int, len(obstacleGrid))
	for i, _ := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	//dp[i][j]表示从0,0到i,j位置的路径数目
	//dp[i][j] = dp[i-1][j] + dp[i][j-1]  if obstacleGrid[i-1][j] == 0  dp[i][j]
	//初始化 没有障碍的情况下
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
