package main

func maxProfit(prices []int) int {
	min := prices[0] //记录当前遍历序列前的最小值
	res := 0         //遍历过程中不断更新结果
	for i := 1; i < len(prices); i++ {
		if res < prices[i]-min {
			res = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return res
}
