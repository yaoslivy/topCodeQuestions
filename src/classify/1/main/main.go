package main

func twoSum(nums []int, target int) []int {
	// 一次遍历，如果当前target-nums[i]值存在于map中则找到两个值和为target，否则将nums[i]的值和下标加入map中
	mm := make(map[int]int, 0)
	for i, val := range nums {
		if index, ok := mm[target-val]; ok {
			return []int{index, i}
		}
		mm[val] = i
	}

	return []int{}
}
