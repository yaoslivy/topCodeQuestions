package main

import (
	"fmt"
	"sort"
)

func main() {
	amount := 5
	coins := []int{1, 2, 5} //[[1 1 1 1 1] [1 1 1 2] [1 2 2] [5]]
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	// dp[j] 凑成金额j的方法数
	// dp[j] += dp[j-coins[i]]  //遍历所有i，得到方法总数
	// dp[0] = 1
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}

var res [][]int

func change2(amount int, coins []int) int {
	if len(coins) == 0 {
		return 0
	}
	res = make([][]int, 0)
	find(coins, amount, []int{})
	sort.Ints(coins)
	//fmt.Println(res)
	return len(res)
}

func find(coins []int, amount int, oneRes []int) {

	sum := 0
	for _, val := range oneRes {
		sum += val
	}
	if sum == amount {
		copyArr := make([]int, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	} else if sum > amount {
		return
	}

	for i := 0; i < len(coins); i++ {
		//重复选取的时候不能选择比集合中最大的元素小的，因为已经在前面选择过了
		if len(oneRes) != 0 && coins[i] < oneRes[len(oneRes)-1] {
			continue
		}
		oneRes = append(oneRes, coins[i])
		find(coins, amount, oneRes)
		oneRes = oneRes[:len(oneRes)-1]
	}
}
