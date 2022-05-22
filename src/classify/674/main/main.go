package main

func findLengthOfLCIS(nums []int) int {
	//dp[i] 以元素nums[i]结尾的连续递增的子序列长度
	//dp[i] = dp[i-1]+1 if nums[i] > nums[i-1]  else dp[i]=1
	//dp[] = 1
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
