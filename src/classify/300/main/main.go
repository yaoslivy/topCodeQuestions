package main

func lengthOfLIS(nums []int) int {
	//dp[i] 以nums[i]为结尾的最长上升子序列长度，dp[i]可以根据dp[j](j<i)推导出来
	//dp[i] = max(dp[i], dp[j]+1) //位置i的最长上升子序列等于j从0到i-1各个位置的最长上升子序列+1的最大值
	//dp[i] = 1
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		//确定一个i位置上的最长递增子序列长度后
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
