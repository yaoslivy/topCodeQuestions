package main

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 局部最优：删除单调坡度上的节点（不包括单调坡度两端的节点），这个坡度就可以有两个局部峰值
	// 整体最优：整个序列有最多的局部峰值，从而达到最长摆动序列
	curDiff, preDiff := 0, 0 //当前一对差值， 前一对差值
	res := 1
	// 2 5
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		//出现峰值，设计巧妙
		if (preDiff <= 0 && curDiff > 0) || (preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}
	return res
}
