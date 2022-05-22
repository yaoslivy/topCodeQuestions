package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res1 := find(nums[1:])           //去掉头
	res2 := find(nums[:len(nums)-1]) //去掉尾
	return max(res1, res2)
}

func find(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	//dp[j] 0-j间房的能取到的最大金额
	//dp[j] = max(dp[j-2]+arr[j], dp[j-1])
	//dp[0] = arr[0], dp[1] = max(arr[0], arr[1])
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])
	for j := 2; j < len(arr); j++ {
		dp[j] = max(dp[j-1], dp[j-2]+arr[j])
	}
	return dp[len(arr)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
