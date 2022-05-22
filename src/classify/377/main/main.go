package main

func combinationSum4(nums []int, target int) int {
	// dp[j] 组合和为j的个数
	// dp[j] += dp[j-nums[i]] //遍历i个数
	// dp[0] = 1
	dp := make([]int, target+1)
	dp[0] = 1
	//先遍历组合值，再遍历候选集合
	//原因是：排列问题，和组合顺序没有关系，这样对于每一个dp[j]，都会经过nums[i-1]和nums[i]的组合，包含了{nums[i-1], nums[i]}和{nums[i], nums[i-1]} 两种情况
	// 先遍历物品，再遍历背包容量则是：先把物品1加入计算，再把物品2加入计算，只会出现{1, 2}，不会有{2, 1}这种情况
	for j := 1; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j-nums[i] >= 0 {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
