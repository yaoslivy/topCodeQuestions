package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(nums)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums) //从小到大排序，目的是：保证枚举的三元满足a<=b<=c

	for a := 0; a < len(nums); a++ { //枚举a
		if a > 0 && nums[a] == nums[a-1] { //a>0是为了保证a-1>=0，如果存在相同的，每次都以相同的元素中第一个为a就可以枚举出不同的情况，不需要以第二个作为a
			continue
		}
		//开始用双指针从两端向中间确定b，c的值
		c := len(nums) - 1
		for b := a + 1; b < c; b++ {
			if b > a+1 && nums[b] == nums[b-1] { //存在相同的b只需要使用第一个b即可
				continue
			}
			for b < c && nums[a]+nums[b]+nums[c] > 0 { //如果三数之和>0说明c值太大
				c--
			}
			if b == c { //在b增加，c减少的过程中相遇，则当前a不满足情况
				break
			}
			if nums[a]+nums[b]+nums[c] == 0 {
				res = append(res, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return res
}
