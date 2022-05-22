package main

import "sort"

var res [][]int

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	// 在有重复元素的序列中，排序的原因是：先选大的元素再选小的元素  == 选小的元素再选了大的元素
	sort.Ints(nums)
	res = make([][]int, 0)
	find(nums, []int{}, 0)
	return res
}

func find(nums []int, oneRes []int, start int) {
	copyArr := make([]int, len(oneRes))
	copy(copyArr, oneRes)
	res = append(res, copyArr)
	//同一层，用过的元素不能再使用了
	mm := make(map[int]struct{})
	for i := start; i < len(nums); i++ {
		if _, ok := mm[nums[i]]; ok { //控制节点start的下一个值不能重复出现
			continue
		}
		mm[nums[i]] = struct{}{}
		oneRes = append(oneRes, nums[i])
		find(nums, oneRes, i+1)
		oneRes = oneRes[:len(oneRes)-1]
	}
}
