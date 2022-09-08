package main

func maxSlidingWindow(nums []int, k int) []int {
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
