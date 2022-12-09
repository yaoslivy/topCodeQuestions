package main

func predictPartyVictory(senate string) string {
	if len(senate) == 0 {
		return ""
	}
	//贪心策略：优先消灭序列后方对方议员，可以使用队列循环记录
	//先使用两个队列记录下两个阵营的议员的下标位置
	RQ := make([]int, 0)
	DQ := make([]int, 0)
	for i, val := range senate {
		if val == 'R' {
			RQ = append(RQ, i)
		} else {
			DQ = append(DQ, i)
		}
	}
	for len(RQ) > 0 && len(DQ) > 0 {
		//同时出队头元素，下标大的元素被消灭，保留下标小的元素
		if RQ[0] < DQ[0] {
			RQ = append(RQ, len(senate)+RQ[0])
		} else {
			DQ = append(DQ, len(senate)+DQ[0])
		}
		RQ = RQ[1:]
		DQ = DQ[1:]
	}
	if len(RQ) > 0 {
		return "Radiant"
	}
	return "Dire"
}
