package main

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	// return smallerNumbersThanCurrentRecursion(nums)

	return smallerNumbersThanCurrentByMap(nums)
}

//map记录+sort
func smallerNumbersThanCurrentByMap(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res := make([]int, len(nums))
	copy(res, nums)
	//直接从小到大排序，每个元素值的位序即是比它小的数字个数
	sort.Ints(res)
	numsMap := make(map[int]int)
	//注意从后往前，这样当两个相同元素时，会记录前一个元素的位序
	for i := len(res) - 1; i >= 0; i-- {
		numsMap[res[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		res[i] = numsMap[nums[i]]
	}
	return res
}

//暴力搜索
func smallerNumbersThanCurrentRecursion(nums []int) []int {
	res := make([]int, len(nums))
	//需确定的位置
	for i := 0; i < len(nums); i++ {
		//比当前位置小的元素个数
		cnt := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i] > nums[j] {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}
