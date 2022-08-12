package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 不需要考虑重复的元素
	// 先用map统计两个数组中元素和出现的次数，再依次从3，4数组中拿出两个元素和map中值求和判断是否为目标值，是则将次数加入结果集

	mm := make(map[int]int, 0) //key为nums1和nums2中两数之和，value放a和b两数和出现的次数
	for _, aVal := range nums1 {
		for _, bVal := range nums2 {
			mm[aVal+bVal]++
		}
	}
	res := 0 //a+b+c+d==0 出现的次数
	for _, cVal := range nums3 {
		for _, dVal := range nums4 {
			if count, ok := mm[-(cVal + dVal)]; ok { //AB + c + d == 0 取出AB个数加入结果
				res += count
			}
		}
	}
	return res
}
