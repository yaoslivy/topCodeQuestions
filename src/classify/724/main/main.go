package main

func pivotIndex(nums []int) int {
	//一次遍历统计总和，然后一次遍历判断减去当前位置的值是否可以使得左边的和为总和的一半
	sum := 0
	for _, val := range nums {
		sum += val
	}
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if sum-nums[i] == leftSum {
			return i
		}
		leftSum += nums[i]
		sum -= nums[i]
	}
	return -1
}
