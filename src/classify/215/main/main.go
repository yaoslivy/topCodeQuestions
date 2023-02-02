package main

func findKthLargest(nums []int, k int) int {
	// 基于快排思路，每次确定中间元素位置下标，如果下标值为k则为目标
	findK(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

func findK(nums []int, left, right, k int) {
	if left > right {
		return
	}
	i, j := left, right
	temp := nums[i]
	for i < j {
		for i < j && temp >= nums[j] {
			j--
		}
		nums[i] = nums[j]
		for i < j && temp <= nums[i] {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = temp
	if i+1 == k {
		return
	}
	findK(nums, left, i-1, k)
	findK(nums, j+1, right, k)
}
