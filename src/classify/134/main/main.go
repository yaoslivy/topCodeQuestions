package main

func canCompleteCircuit(gas []int, cost []int) int {
	//局部最优：当前累加剩余油量和curSum一旦小于0，则起始位置至少是j+1
	//全局最优：找到可以跑一圈的起始位置
	// 如果总油量-总消耗 >= 0 那么一定可以跑完一圈，否则不能
	//若要跑完一圈，则加的全部油量总和 > 消耗的全部油量总和，且每一段路消耗的油量都必须小于剩余油量
	restSum := 0
	curSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		restSum += gas[i]
		restSum -= cost[i]
		// 当前剩余油量需要能维持这段路程
		curSum += gas[i]
		curSum -= cost[i]
		if curSum < 0 { //如果以start开始不能维持其中的一段路程
			curSum = 0
			start = i + 1
		}
	}
	if restSum < 0 {
		return -1
	}
	return start
}

func canCompleteCircuit2(gas []int, cost []int) int {
	// 暴力求解：遍历每一个加油站为起点的情况，模拟一圈
	for i := 0; i < len(gas); i++ { //起点
		// costGas := cost[i] //到i+1需要的消耗
		curI := i //当前位置
		curGas := gas[curI]
		for curGas >= cost[curI] {
			curGas -= cost[curI]         //减去当前消耗
			curI = (curI + 1) % len(gas) //下一个位置
			if curI == i {
				return i
			}
			curGas += gas[curI] //加油
		}
	}
	return -1
}
