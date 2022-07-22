package main

func jump2(nums []int) int {
	// 在当前位置所能到达的范围必须跳一次，然后不断更新在这个区间内所能到达的最大右边界
	if len(nums) <= 1 {
		return 0
	}
	curRight := nums[0] //记录当前所能到达的右边界
	maxRight := nums[0] //记录一次跳跃区间内所能到达的下一个最大右边界
	res := 1
	for i := 1; i <= maxRight; i++ {
		maxRight = max(maxRight, i+nums[i])
		if curRight >= len(nums)-1 {
			break
		}
		if i == curRight { //如果已经到达该区间的最大范围，更新下一个跳跃区间
			curRight = maxRight
			res++
		}
	}
	return res
}

func jump(nums []int) int {
	// 局部最优：当前覆盖范围遍历，不断更新下一个最大覆盖范围
	// 全局最优：一步尽可能多走，从而达到最小步数
	if len(nums) == 1 {
		return 0
	}
	curDis := 0 //当前覆盖最远距离下标
	res := 0    //记录走的最大步数
	right := 0  // 下一步覆盖最远距离下标
	for i := 0; i <= right; i++ {
		right = max(i+nums[i], right) //不断更新下一步所能覆盖最远距离下标
		if curDis >= len(nums)-1 {    //如果当前可以跳到最后位置
			break
		}
		if i == curDis { //i移动到了当前最远覆盖范围
			res++
			curDis = right
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
