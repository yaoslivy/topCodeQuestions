package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	// 先从小到大排序
	sort.Ints(nums)
	res := make([][]int, 0)
	// 先固定a, 然后固定b, 然后用c，d分别指向右端剩余序列的左右端点，根据求和不断缩短区间，注意对a,b,c,d的去重
	for a := 0; a < len(nums)-3; a++ {
		//如果nums[a] > target && target >= 0，则剩余序列的求和不可能等于target
		if nums[a] > target && target >= 0 {
			break
		}
		// 对a去重
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			// 对b去重
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			c, d := b+1, len(nums)-1
			for c < d {
				if nums[a]+nums[b]+nums[c]+nums[d] > target {
					d--
				} else if nums[a]+nums[b]+nums[c]+nums[d] < target {
					c++
				} else {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
					// 对c，d去重
					for c < d && nums[c] == nums[c-1] {
						c++
					}
					for c < d && nums[d] == nums[d+1] {
						d--
					}
				}
			}
		}
	}
	return res
}
