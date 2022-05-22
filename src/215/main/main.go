package main

import "sort"

func findKthLargest(nums []int, k int) int {
	k-- //数组元素从0开始
	// findKByQuickSort(nums, 0, len(nums)-1, k)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k]
}

func findKByQuickSort(arr []int, left, right, k int) {
	i := left
	j := right
	temp := arr[i]
	for i < j {
		for i < j && temp >= arr[j] { //尾指针指向的元素比temp大，才换到i位置,每次将尾部大的往前放，将首部小的往后放
			j--
		}
		arr[i] = arr[j]
		for i < j && temp < arr[i] {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = temp
	if i == k {
		return
	}
	if left < i && k < i { //k在i左边才需要递归找
		findKByQuickSort(arr, left, i-1, k)
	}
	if i < right && k > i { //k在i右边才需要递归找
		findKByQuickSort(arr, i+1, right, k)
	}
}
