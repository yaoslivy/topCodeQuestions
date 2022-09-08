package main

import (
	"math"
	"strconv"
)

func main() {

}

//基数排序
func radixSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	//先找到arr中最大的元素
	maxVal := arr[0]
	minVal := arr[0] //记录arr中最小值，如果最小值为负数则需要把所有负数都变为正数
	for _, val := range arr {
		maxVal = max(maxVal, val)
		minVal = min(minVal, val)
	}
	//将所有值都加上负数的绝对值
	if minVal < 0 {
		maxVal -= minVal
		for i := 0; i < len(arr); i++ {
			arr[i] -= minVal
		}
	}
	//当前最大值的长度
	maxValLen := len(strconv.Itoa(maxVal))
	//设置10个桶，表示0到9，每个桶的长度最大为len(arr)
	radixNums := make([][]int, 10)
	for i := 0; i < 10; i++ {
		radixNums[i] = make([]int, len(arr))
	}
	radixCount := make([]int, 10)    //记录当前桶中的元素个数
	for k := 0; k < maxValLen; k++ { //需要排序的位次
		for _, val := range arr {
			temp := val / int(math.Pow(10, float64(k))) % 10 //求到当前元素位次的值
			radixNums[temp][radixCount[temp]] = val
			radixCount[temp]++
		}
		//将元素按照次序收集一遍
		index := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < radixCount[i]; j++ {
				arr[index] = radixNums[i][j]
				index++
			}
			radixCount[i] = 0
		}
	}

	//最后减去负数的绝对值
	if minVal < 0 {
		for i := 0; i < len(arr); i++ {
			arr[i] += minVal
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
