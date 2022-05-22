package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// 等价于求两端数字中公共子序列的最大长度
	// dp[i][j] 表示0至i-1位置和0至j-1位置公共子序列的最大长度
	// dp[i][j] = dp[i-1][j-1] if nums1[i-1] == nums2[j-1]
	// else  = max(dp[i-1][j], dp[i][j-1]) //去掉i-1位置； 去掉j-1位置
	// dp[i][0] = 0 dp[0][j] = 0
	dp := make([][]int, len(nums1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
