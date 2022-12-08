package main

func rotate(nums []int, k int) {
	if k <= 0 {
		return
	}
	// 1,2,3,4,5,6,7
	// 4,3,2,1,7,6,5
	// 5,6,7,1,2,3,4
	k = k % len(nums)
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
