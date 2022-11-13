package main

import "sort"

func nextPermutation(nums []int) {
	//从后往前遍历，
	//尽量将后面小的数替换到前面比其略大的位置，如此使得替换后的数字略大
	if len(nums) <= 1 {
		return
	}

	for i := len(nums) - 2; i >= 0; i-- { //左边比右边略大的值
		//在i值的右边，从后往前找到一个比其小的值，找到后替换，然后将替换之后位置的值进行从小到大排序
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				sort.Ints(nums[i+1:])
				return
			}
		}
	}
	sort.Ints(nums)
	return
}
