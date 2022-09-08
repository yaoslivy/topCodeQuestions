package main

func removeElement(nums []int, val int) int {
	//使用双指针，left指针处数值等于val，则和后面不为val的元素互换位置
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	res := 0
	for left <= right {
		if nums[left] == val {
			//从后往前找到一个不为val的值
			for left <= right && nums[right] == val {
				right--
			}
			//没有不为val的的值了
			if left > right {
				break
			}
			//交换
			nums[left], nums[right] = nums[right], nums[left]
		} else {
			res++
			left++
		}
	}
	return res
}
