package main

func maxSlidingWindow(nums []int, k int) []int {
	//使用单调队列保存访问过的元素下标，队头到队尾对应元素值从大到小
	//当当前元素值大于队尾元素时，一直将队尾元素出队，直到小于等于再入队
	if len(nums) == 0 {
		return []int{}
	}
	Q := make([]int, 0)
	res := make([]int, 0)
	//先处理前k个元素
	for i := 0; i < k; i++ {
		//当前元素>队尾元素，出队
		for len(Q) != 0 && nums[i] > nums[Q[len(Q)-1]] {
			Q = Q[:len(Q)-1]
		}
		//队列为空或当前元素值<=队尾元素值
		Q = append(Q, i)
	}
	//第一个窗口最大值
	res = append(res, nums[Q[0]])
	//处理剩余元素
	for i := k; i < len(nums); i++ {
		//窗口向右滑动了一个位置，判断队头元素是否需要出队
		if Q[0] == i-k {
			Q = Q[1:]
		}
		//当前元素>队尾元素，出队
		for len(Q) != 0 && nums[i] > nums[Q[len(Q)-1]] {
			Q = Q[:len(Q)-1]
		}
		//队列为空或当前元素值<=队尾元素值
		Q = append(Q, i)
		//记录当前窗口最大值
		res = append(res, nums[Q[0]])
	}
	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// 维护一个单调递减的队列，当新来一个元素小于当前队列队尾部元素则加入，
	Q := make([]int, 0)
	res := make([]int, 0)
	// 1,3,-1,-3
	//先将判断前k个元素
	for i := 0; i < k; i++ {
		//新来的元素比队列末尾元素大则将队列末尾元素出队
		for len(Q) != 0 && Q[len(Q)-1] < nums[i] {
			Q = Q[:len(Q)-1]
		}
		Q = append(Q, nums[i])
	}
	res = append(res, Q[0])
	//处理k后面元素
	for i := k; i < len(nums); i++ {
		//移除滑动窗口左边界元素，如果左边界元素为队头元素值，则移除
		if nums[i-k] == Q[0] {
			Q = Q[1:]
		}
		//新来的元素比队列末尾元素大则将队列末尾元素出队
		for len(Q) != 0 && Q[len(Q)-1] < nums[i] {
			Q = Q[:len(Q)-1]
		}
		Q = append(Q, nums[i])
		res = append(res, Q[0])
	}
	return res
}
