package main

import "fmt"

//有N件物品和一个最多能被重量为W 的背包。第i件物品的重量是weight[i],得到的价值是value[i] 。
//每件物品只能用一次,求解将哪些物品装入背包里物品价值总和最大。
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
	for _, val := range dp {
		fmt.Println(val)
	}
	dp2 := findMaxVal2(weight, value, volume)
	fmt.Println("----------------")
	fmt.Println(dp2)
}

func findMaxVal2(weight []int, value []int, volume int) []int {
	//dp[j]为容量为j的背包，所背的物品价值最大为dp[j]
	//dp[j] = dp[j], dp[j-weight[i]] + value[i]
	dp := make([]int, volume+1)
	for i := 0; i < len(weight); i++ { //遍历物品
		for j := volume; j >= weight[i]; j-- { //遍历背包容量; 倒序遍历为了保证物品i只被放入一次
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	//for j := volume; j > 0; j-- { //遍历背包容量; 倒序遍历为了保证物品i只被放入一次
	//	for i := 0; i < len(weight); i++ { //遍历物品
	//		fmt.Println("weight[i]=", weight[i])
	//		fmt.Println("j-weight[i]=", j-weight[i])
	//		if j-weight[i] >= 0 {
	//			fmt.Println("dp[j-weight[i]]=", dp[j-weight[i]])
	//			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
	//			fmt.Println("dp[j]=", dp[j])
	//		}
	//	}
	//	fmt.Println()
	//}
	return dp
}

func findMaxVal(weight []int, value []int, volume int) [][]int {
	// dp[i][j]表示从下标0-i的物品中任意取，放进容量为j的背包中的价值总和
	// dp[i][j] = dp[i-1][j], dp[i-1][j-weight[i]]+value[i] //容量为j，不放入物品i则价值就是i-1件物品的价值总和，放入物品i的价值总和为物品i的价值+容量-value[i]的价值总和
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, volume+1)
	}
	// 初始化
	for i := 1; i <= volume; i++ {
		dp[0][i] = value[0]
	}
	//先遍历物品，再遍历背包容量
	for i := 1; i < len(weight); i++ { //遍历物品
		for j := 0; j <= volume; j++ { //遍历背包容量
			if j-weight[i] >= 0 { //取下标0-i位置的物品
				dp[i][j] = max(dp[i-1][j-weight[i]]+value[i], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	//先遍历背包容量，再遍历遍历物品
	//for j := 0; j <= volume; j++ {
	//	for i := 1; i < len(weight); i++ {
	//		if j-weight[i] >= 0 { //取下标0-i位置的物品
	//			dp[i][j] = max(dp[i-1][j-weight[i]]+value[i], dp[i-1][j])
	//		} else {
	//			dp[i][j] = dp[i-1][j]
	//		}
	//	}
	//}
	return dp
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
