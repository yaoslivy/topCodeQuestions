package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 二分搜索，每次二分搜索，一定一边是有序的，根据有序的一边来判断target是否在其中
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		//左边有序
		if nums[0] <= nums[mid] {
			// target在前一段
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// target在后一段
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
