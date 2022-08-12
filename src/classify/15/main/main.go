package main

import "sort"

func threeSum(nums []int) [][]int {
	//先从小到大排序
	sort.Ints(nums)
	// 固定a值，剩余右边区间设置b,c分别指向两端，然后求和，> 0 则c--， < 0 则b++，注意对a，b，c的去重
	// -1,0,1,2,-1,-4
	// -4,-1,-1,0,1,2
	res := make([][]int, 0)
	for a := 0; a < len(nums)-2; a++ {
		if nums[a] > 0 { //排序后，b,c位置都是>0的，求和后不可能=0
			break
		}
		//当前a值重复，跳过
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b, c := a+1, len(nums)-1
		for b < c {
			if nums[a]+nums[b]+nums[c] > 0 {
				c--
			} else if nums[a]+nums[b]+nums[c] < 0 {
				b++
			} else {
				res = append(res, []int{nums[a], nums[b], nums[c]})
				b++
				c--
				//对b，c去重
				for b < c && nums[b] == nums[b-1] {
					b++
				}
				for b < c && nums[c] == nums[c+1] {
					c--
				}
			}
		}
	}
	return res
}
