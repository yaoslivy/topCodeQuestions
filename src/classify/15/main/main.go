package main

import "sort"

func threeSum(nums []int) [][]int {
	//先将nums按照从小到大排序
	sort.Ints(nums)
	//枚举i, j, k 处下标的值
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { //有序序列的第一个值大于0，则无法求和为0
			break
		}
		//对i处元素去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//枚举j,k值，从剩余序列的两端向中间逼近
		j, k := i+1, len(nums)-1
		for j < k {
			//对j处元素去重
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			//对k处元素去重
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
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
