package main

import "sort"

var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	res = make([][]int, 0)
	find(candidates, target, []int{}, 0)
	return res
}

func find(arr []int, target int, oneRes []int, start int) {
	sum := 0
	for _, val := range oneRes {
		sum += val
	}
	if sum > target { //排序后，后面的值都会大于target，不需要再添加值了
		return
	}
	if sum == target {
		copyArr := make([]int, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}

	mm := make(map[int]struct{}) // 一个元素值在同一位置只能使用一次
	for i := start; i < len(arr); i++ {
		if _, ok := mm[arr[i]]; ok {
			continue
		}
		mm[arr[i]] = struct{}{}
		oneRes = append(oneRes, arr[i])
		find(arr, target, oneRes, i+1)
		oneRes = oneRes[:len(oneRes)-1]
	}
}
