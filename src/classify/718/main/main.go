package main

func findLength(nums1 []int, nums2 []int) int {
	dp := make([]int, len(nums1)+1)
	res := 0
	for i := 1; i <= len(nums1); i++ {
		for j := len(nums2); j > 0; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else { //
				dp[j] = 0
			}
			res = max(res, dp[j])
		}
	}
	return res
}

func findLength2(nums1 []int, nums2 []int) int {
	// dp[i][j] 表示以nums1[i-1]结尾和以nums2[j-1]结尾的序列中最长重复子数组长度
	// dp[i][j] = dp[i-1][j-1]+1 if nums[i-1] == nums[j-1] else = 0
	// dp[i][0] = dp[0][j] = 0
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(nums1)+1; i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	res := 0
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
