package main

func jump(nums []int) int {
	// 局部最优：当前覆盖范围遍历，不断更新下一个最大覆盖范围
	// 全局最优：一步尽可能多走，从而达到最小步数
	if len(nums) == 1 {
		return 0
	}
	curDis := 0  //当前覆盖最远距离下标
	res := 0     //记录走的最大步数
	nextDis := 0 // 下一步覆盖最远距离下标
	for i := 0; i < len(nums); i++ {
		nextDis = max(i+nums[i], nextDis) //不断更新下一步所能覆盖最远距离下标
		if curDis >= len(nums)-1 {        //如果当前可以跳到最后位置
			break
		}
		if i == curDis { //i移动到了当前最远覆盖范围
			res++
			curDis = nextDis
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
