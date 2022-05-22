package main

import "fmt"

func main() {
	nums := []int{9, 7, 0, 3, 9, 8, 6, 5, 7, 6}
	fmt.Println(findTargetSumWays(nums, 2))
}

func findTargetSumWays(nums []int, target int) int {
	//正数集合+负数集合=sum
	//正数集合-负数集合=target
	//正数集合 = (sum+target)/2    (sum+target) >= 0   (sum+target) % 2 == 0
	// 等价于在数字集合中找到和(sum+target)/2的数
	// dp[j] 找到和为j的集合的组合数
	// dp[j] += dp[j-nums[i]]  //0至i-1件物品中和凑成j的方法数 + 使用第i件物品凑成的方法数为凑成的j-nums[i]
	// dp[0] = 1
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if (sum+target) < 0 || (sum+target)%2 != 0 {
		return 0
	}
	aim := (sum + target) / 2
	dp := make([]int, aim+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := aim; j-nums[i] >= 0; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[aim]
}
