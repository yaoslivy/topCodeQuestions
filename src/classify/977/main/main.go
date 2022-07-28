package main

func sortedSquares(nums []int) []int {
	//递增序列，平方后最大值一定在两边去的，设置双指针观察平方后最大值取值
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	k := len(nums) - 1
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] { //头指针指向的值平方后更大
			res[k] = nums[i] * nums[i]
			k--
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
			k--
		}
	}
	return res
}
