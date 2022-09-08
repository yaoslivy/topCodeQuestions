package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// return nextGreaterElementTraverse(nums1, nums2)
	return nextGreaterElementByStack(nums1, nums2)
}

// 使用单调栈
func nextGreaterElementByStack(nums1, nums2 []int) []int {
	//先使用map记录nums1中所有元素值出现的下标，便于通过值迅速找到下标
	mm := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		mm[nums1[i]] = i
	}
	//使用单调栈记录访问过的元素，保持从栈底到栈顶递减，
	//当下一个元素大于栈顶元素时，则说明栈顶元素找到了下一个最大值，然后一直出栈，不断判断当前访问元素和栈顶元素的大小
	S := make([]int, 0)
	res := make([]int, len(nums1))
	for i := 0; i < len(res); i++ {
		res[i] = -1 //找不到则默认值为-1
	}
	//需要确定的个数，最后一位一定是-1
	for i := 0; i < len(nums2); i++ {
		//当前元素比栈顶元素值大，则栈顶元素找到了后面更大的元素
		for len(S) != 0 && nums2[i] > nums2[S[len(S)-1]] {
			if index, ok := mm[nums2[S[len(S)-1]]]; ok {
				res[index] = nums2[i]
			}
			S = S[:len(S)-1]
		}
		//栈为空或当前元素比栈顶元素值小
		S = append(S, i)
	}
	return res
}

//暴力求解
func nextGreaterElementTraverse(nums1, nums2 []int) []int {
	res := make([]int, len(nums1))
	//需要确定的位置
	for i := 0; i < len(nums1); i++ {
		//在nums2中找到对应位置右侧的第一个比i位置元素大的元素
		flag := false // 是否找到该位置
		j := 0
		for ; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				flag = true
			}
			if flag == true && nums2[j] > nums1[i] {
				res[i] = nums2[j]
				break
			}
		}
		if j == len(nums2) { //当前元素之后未有比当前元素值的元素
			res[i] = -1
		}
	}
	return res
}
