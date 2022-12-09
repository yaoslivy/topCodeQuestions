package main

func moveZeroes(nums []int) {
	moveZeroesDoublePointer(nums)
	// moveZeroesInsert(nums)
}

//双指针思路
func moveZeroesDoublePointer(nums []int) {
	if len(nums) <= 1 {
		return
	}
	//设置双指针同时从0开始移动，当left遇到0时，right找到一个非0的元素插入left
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 { //如果当前right位置元素不为0，则插入left位置，为0则right多走一步
			nums[left] = nums[right]
			left++
		}
	}
	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroesDoublePointer2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	//使用双指针,left right 同时从移动
	//当right处值不为0时，将right处值插入left位置， left，right同时移动
	//当right处值为0时，将right向后移动，left不移动
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	//left之后已经没有非0元素
	for left < len(nums) {
		nums[left] = 0
		left++
	}
}

//直接插入排序思路
func moveZeroesInsert(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		j := i - 1
		temp := nums[j+1]
		//找到一个非0的位置，插入
		for j >= 0 && nums[j] == 0 {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
}
