package main

func intersection(nums1 []int, nums2 []int) []int {
	//先遍历nums1，遍历过程中使用map记录nums1中每个元素出现的情况，
	//再遍历nums2，判断每一个元素是否在map中出现，出现则添加到结果集中，同时将map中对应元素置false，防止重复
	mm := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		mm[nums1[i]] = struct{}{}
	}
	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if _, ok := mm[nums2[i]]; ok {
			res = append(res, nums2[i])
			delete(mm, nums2[i])
		}
	}

	return res
}
