package main

var res [][]int

func findSubsequences(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res = make([][]int, 0)
	find(nums, []int{}, 0)
	return res
}

func find(nums []int, oneRes []int, start int) {
	if len(oneRes) > 1 {
		if oneRes[len(oneRes)-1] >= oneRes[len(oneRes)-2] {
			copyArr := make([]int, len(oneRes))
			copy(copyArr, oneRes)
			res = append(res, copyArr)
		} else {
			return
		}
	}
	//使用set记录该层使用过的元素，保证不重复，只在本层使用
	mm := make(map[int]struct{})
	for i := start; i < len(nums); i++ {
		if _, ok := mm[nums[i]]; ok { //如果该位置选择了重复元素，则跳过
			continue
		}
		mm[nums[i]] = struct{}{}
		oneRes = append(oneRes, nums[i])
		find(nums, oneRes, i+1)
		oneRes = oneRes[:len(oneRes)-1]
	}
}
