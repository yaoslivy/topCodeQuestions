package main

import "fmt"

func main() {
	dp := make([][4]int, 10)
	fmt.Printf("%T", dp[0])
	arr := [2]int{1, 2}
	fmt.Printf("%T", arr)
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		copyArr := make([]int, len(arr))
		merge(arr, copyArr, left, mid, right)
	}
}

func merge(arr, copyArr []int, left, mid, right int) {
	i, j := left, mid+1
	index := 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			copyArr[index] = arr[i]
			i++
			index++
		} else {
			copyArr[index] = arr[j]
			j++
			index++
		}
	}
	for i <= mid {
		copyArr[index] = arr[i]
		i++
		index++
	}
	for j <= right {
		copyArr[index] = arr[j]
		j++
		index++
	}
	index = 0
	leftIndex := left
	for leftIndex <= right {
		arr[leftIndex] = copyArr[index]
		index++
		leftIndex++
	}
}
