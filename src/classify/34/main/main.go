package main

func searchRange(nums []int, target int) []int {
	aimIndex := biSearch(nums, target) //找到目标值在数组中的一个位置
	if aimIndex == -1 {                //找不到目标值
		return []int{-1, -1}
	}
	res := make([]int, 2)
	index := aimIndex
	//向右找目标值的最大下标
	for index < len(nums) && nums[index] == target {
		res[1] = index
		index++
	}
	index = aimIndex
	//向左找目标值的最小下标
	for index >= 0 && nums[index] == target {
		res[0] = index
		index--
	}

	return res
}

//二分查找目标元素在一段排序数组中的位置，目标元素可能有多个，返回其中的一个下标
func biSearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
