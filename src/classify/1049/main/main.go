package main

func lastStoneWeightII(stones []int) int {
	//尽可能的将石头均分  sum/2  求最接近均分方案的石头组合和
	//dp[j] 最接近j的石头组合和
	//dp[j] = max(dp[j], dp[j-stones[i]]+stones[i]) //目标j 从1至i-1个石头选择的最大组合和 选择第i个石头所能获取的最大组合和=stones[i] + dp[j-stones[i]]
	//dp[0] = 0
	sum := 0
	for _, val := range stones {
		sum += val
	}
	aim := sum / 2 //向下取整
	dp := make([]int, aim+1)
	for i := 0; i < len(stones); i++ {
		for j := aim; j-stones[i] >= 0; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	//dp[aim] 较小的组合和 sum-dp[aim] 较大的组合和
	return sum - dp[aim]*2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
