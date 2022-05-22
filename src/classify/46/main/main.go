package main

var res [][]int

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res = make([][]int, 0)
	mm := make(map[int]struct{}) //记录当前层使用的数字，下一层不能使用
	find(nums, []int{}, mm)
	return res
}

func find(nums []int, oneRes []int, mm map[int]struct{}) {
	if len(oneRes) == len(nums) {
		copyArr := make([]int, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}

	for i := 0; i < len(nums); i++ { //遍历当前位置可能的取值
		if _, ok := mm[nums[i]]; ok { //当前值在上层已经使用了，也就是说在递归路线中不能再使用了
			continue
		}
		mm[nums[i]] = struct{}{}
		oneRes = append(oneRes, nums[i])
		find(nums, oneRes, mm)
		oneRes = oneRes[:len(oneRes)-1]
		delete(mm, nums[i]) //递归路线走完后需要删除该值，
	}
}
