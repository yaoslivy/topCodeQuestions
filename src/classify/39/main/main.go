package main

import "sort"

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	res = make([][]int, 0)
	find(candidates, []int{}, target)
	return res
}

func find(candidates []int, oneRes []int, target int) {
	sum := 0
	for _, val := range oneRes {
		sum += val
	}
	if sum == target {
		copyArr := make([]int, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}
	if sum > target {
		return
	}
	for i := 0; i < len(candidates); i++ {
		//重复选取的时候不能选择比之前元素小的值，因为已经选择过了
		if len(oneRes) != 0 && candidates[i] < oneRes[len(oneRes)-1] {
			continue
		}
		oneRes = append(oneRes, candidates[i])
		find(candidates, oneRes, target)
		oneRes = oneRes[:len(oneRes)-1]
	}

}

func find2(candidates []int, target int, oneRes []int, start int) {
	sum := 0
	for _, val := range oneRes {
		sum += val
	}
	if sum == target {
		copyArr := make([]int, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}
	if sum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		oneRes = append(oneRes, candidates[i])
		//每次递归选择时从当前位置之后的元素开始选择，因为当前位置可以重复选择,
		find2(candidates, target, oneRes, i)
		oneRes = oneRes[:len(oneRes)-1]
	}
}
