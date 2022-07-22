package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Println("this is a=", a)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	curline := strings.Split(sc.Text(), " ")
	arr := make([]int, a)
	fmt.Println(curline)
	for i := 0; i < a; i++ {
		num, _ := strconv.Atoi(curline[i])
		arr[i] = num
	}
	fmt.Println(arr)
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
