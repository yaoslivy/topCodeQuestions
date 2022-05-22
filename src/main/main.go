package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	intn := rand.Intn(104)
	fmt.Println(intn)
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
