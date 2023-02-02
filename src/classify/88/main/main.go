package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 归并思路
	// 逆序方式从后往前填充
	if n == 0 {
		return
	}
	i, j := m-1, n-1
	index := len(nums1) - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
	for i >= 0 {
		nums1[index] = nums1[i]
		index--
		i--
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}
}
