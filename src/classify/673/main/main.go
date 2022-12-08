package main

func findNumberOfLIS(nums []int) int {
	//在统计子序列的长度时，同时统计该个数
	//dp[i] 表示以nums[i]元素结尾的子序列最长长度
	//count[i] 表示以nums[i]元素结尾的子序列长度的个数
	//dp[i] = max(dp[i], dp[j]+1) if nums[i] > nums[j]
	//count[i] = count[j] if dp[j]+1 > dp[i]
	//  = count[i] + count[j] if dp[j]+1 == dp[i] //表示中间元素的替换，最长子序列的长度保持不变
	//dp[i] = 1
	//count[i] = 1
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
	}
	maxLen := 0 //记录最长递增子序列的长度
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
				dp[i] = max(dp[i], dp[j]+1)
			}
			maxLen = max(maxLen, dp[i])
		}
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if maxLen == dp[i] {
			res += count[i]
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
