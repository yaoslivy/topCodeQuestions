package main

func candy(ratings []int) int {
	// 局部最优：只要右边评分大于左边，右边的孩子就多一个糖果，
	// 全局最优：相邻的孩子中，评分高的右孩子获得比左边孩子更多的糖果

	arr := make([]int, len(ratings))
	for i, _ := range arr {
		arr[i] = 1
	}
	//从前往后，只要右边的大于左边，右边就在左边的基础上+1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		}
	}
	//从后往前，
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			arr[i] = max(arr[i], arr[i+1]+1) //既要>arr[i-1] 又要>arr[i+1]
		}
	}
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
