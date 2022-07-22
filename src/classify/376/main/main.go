package main

func wiggleMaxLength2(nums []int) int {
	// 贪心策略： 最佳的摆动序列是一段序列的两个端点值分别为当前范围的最大值和最小值
	// 则记录前一段序列的单调性和后一段序列的单调性，遇到后一段相反序列则res+1，更新前一段序列的单调性
	if len(nums) == 0 {
		return 0
	}
	curDiff := 0
	preDiff := 0
	res := 1 //只要存在一个元素就有一个摆动序列
	for i := 1; i < len(nums); i++ {
		curDiff = nums[i] - nums[i-1]
		//遇到相反序列
		if (preDiff <= 0 && curDiff > 0) || (preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}
	return res
}


func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 局部最优：删除单调坡度上的节点（不包括单调坡度两端的节点），
	// 这个坡度就可以有两个局部峰值
	// 整体最优：整个序列有最多的局部峰值，从而达到最长摆动序列
	curDiff, preDiff := 0, 0 //前一段单调序列上的差值， 后一段单调序列上的差值

	res := 1
	// 2 5
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		//当两差值一正一负则可以开始统计下一段单调序列了
		if (preDiff <= 0 && curDiff > 0) || (preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}
	return res
}
