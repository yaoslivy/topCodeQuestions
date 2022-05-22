package main

import "fmt"

//有N件物品和一个最多能背重量为W 的背包。第i件物品的重量是weight[i],得到的价值是value[i] 。
//每件物品能用多次,求解将哪些物品装入背包里物品价值总和最大。
//背包的最大容量为4
//			重量			价值
// 物品0  	1			15
// 物品1  	3			20
// 物品2  	4			30

func main() {
	//N件物品， 每个物品的重量不同，每个物品对应不同的价值
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	volume := 4
	dp := findMaxVal(weight, value, volume)
	fmt.Println(dp)
	dp2 := findMaxVal2(weight, value, volume)
	fmt.Println(dp2)
}

func findMaxVal(weight []int, value []int, volume int) []int {
	dp := make([]int, volume+1)
	for i := 0; i < len(weight); i++ { //遍历物品
		for j := weight[i]; j <= volume; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp
}

func findMaxVal2(weight []int, value []int, volume int) []int {
	dp := make([]int, volume+1)

	for j := 0; j <= volume; j++ {
		for i := 0; i < len(weight); i++ { //遍历物品
			if j-weight[i] >= 0 {
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
		}
	}
	return dp
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
