package main

func sortArrayByParityII(nums []int) []int {
	//双指针，一个指向奇数位置，一个指向偶数位置，找到不符合的值交换
	left, right := 0, 1
	for left < len(nums)-1 && right < len(nums) {
		//找到left不符合的值
		for left < len(nums)-1 && nums[left]%2 == 0 {
			left += 2
		}
		//找到right不符合的值
		for right < len(nums) && nums[right]%2 == 1 {
			right += 2
		}
		if left < len(nums)-1 && right < len(nums) {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return nums
}
