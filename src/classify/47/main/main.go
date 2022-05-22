package main

import "sort"

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	visited := make([]bool, len(nums))
	// sort.Ints(nums)
	find(nums, []int{}, visited)
	return res
}

func find(nums, oneRes []int, visited []bool) {
	if len(oneRes) == len(nums) {
		copyArr := make([]int, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}
	mm := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		//i位置元素值的下一个值不能重复出现，同一位置上的元素在一段深搜序列中不能重复出现
		if _, ok := mm[nums[i]]; ok || visited[i] == true {
			continue
		}
		mm[nums[i]] = struct{}{}
		visited[i] = true
		oneRes = append(oneRes, nums[i])
		find(nums, oneRes, visited)
		visited[i] = false
		oneRes = oneRes[:len(oneRes)-1]
	}
}

func permuteUnique2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums) //排序便于通过相邻节点来判断是否该位置重复使用元素值相同的元素了
	res = make([][]int, 0)
	visited := make([]bool, len(nums))
	find2(nums, []int{}, visited)
	return res
}

func find2(nums []int, oneRes []int, visited []bool) {
	if len(oneRes) == len(nums) {
		copyArr := make([]int, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == true { //该节点在该位置使用过，一个元素不能重复使用
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && visited[i-1] == true { //如何当前位置已经使用了该元素，则不能再选择和该元素值相同的元素
			continue
		}
		visited[i] = true
		oneRes = append(oneRes, nums[i])
		find(nums, oneRes, visited)
		oneRes = oneRes[:len(oneRes)-1]
		visited[i] = false

	}
}
