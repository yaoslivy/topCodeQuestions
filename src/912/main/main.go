package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	arr := []int{3, 2, 5, 1, 4, 12, 32, 23, 7, 9, 5}
	fmt.Println(arr)
	sortArray(arr)
	fmt.Println(arr)
}

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(arr []int, left, right int) {
	i := left
	j := right
	rand.Seed(time.Now().Unix())
	randNums := math.Abs(float64(right - left))
	randIndex := left
	if randNums != 0 {
		randIndex = left + rand.Intn(int(randNums))
	}
	temp := arr[randIndex]
	for i < j {
		for i < j && temp <= arr[j] { //从后往前找一个小于temp的数来交换
			j--
		}
		arr[i] = arr[j]
		for i < j && temp > arr[i] { // 从前往后找一个大于等于temp的数来交换
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = temp //i==j
	if left < i {
		quickSort(arr, left, i-1)
	}
	if j < right {
		quickSort(arr, j+1, right)
	}
}
