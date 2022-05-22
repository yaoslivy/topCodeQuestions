package main

var res [][]int

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res = make([][]int, 0)
	find(nums, 0, []int{})
	return res
}

func find(nums []int, start int, onRes []int) {
	copyOneRes := make([]int, len(onRes))
	copy(copyOneRes, onRes)
	res = append(res, copyOneRes)

	for i := start; i < len(nums); i++ { //递归添加
		onRes = append(onRes, nums[i])
		find(nums, i+1, onRes)
		onRes = onRes[:len(onRes)-1] //回溯
	}
}
