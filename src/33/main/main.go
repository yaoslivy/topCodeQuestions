package main

func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		//mid左右两边，一定有一边有序，根据有序的一边来判断target值是否在其中
		if nums[0] <= nums[mid] { //前一段有序
			if nums[0] <= target && target < nums[mid] { //target值在前一段
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { //后一段有序
			if nums[mid] < target && target <= nums[len(nums)-1] { //target值在后一段
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	j := len(nums) - 1
	for j >= 0 {
		if nums[j] == target {
			return j
		}
		if j > 0 && nums[j-1] > nums[j] {
			j--
			break
		}
		j--
	}
	if j != 0 {
		return binSearch(nums, 0, j, target)
	} else if nums[0] == target {
		return 0
	} else {
		return -1
	}
}

func binSearch(arr []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) >> 1
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right--
		} else if arr[mid] < target {
			left++
		}
	}
	return -1
}
