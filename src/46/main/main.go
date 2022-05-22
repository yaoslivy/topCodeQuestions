package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	ans := permute(nums)
	for _, val := range ans {
		fmt.Print(val, " ")
	}
}

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	backtrack(nums, []int{})
	return res
}

func backtrack(nums []int, path []int) {
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) //每次移除i位置元素，直接使用切片，将nums中i+1位置后的元素解包添加到nums中i位置之后
		backtrack(nums, path)
		path = path[:len(path)-1]                                   //移除添加在末尾的元素
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //将i位置元素加回来，元素位置不能变
	}
}
