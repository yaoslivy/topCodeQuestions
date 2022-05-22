package main

var res [][]int

func combinationSum3(k int, n int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	res = make([][]int, 0)
	find(n, k, 1, []int{})
	return res
}

func find(n, k int, start int, oneRes []int) {
	if len(oneRes) == k {
		sum := 0
		for _, val := range oneRes {
			sum += val
		}
		if sum == n {
			copyArr := make([]int, k)
			copy(copyArr, oneRes)
			res = append(res, copyArr)
		}
		return
	}
	//还需要添加的元素长度: k-len(oneRes) ，当集合中元素小于这个长度了便不可能是一个答案
	for i := start; i <= 9-(k-len(oneRes))+1; i++ {
		oneRes = append(oneRes, i)
		find(n, k, i+1, oneRes)
		oneRes = oneRes[:len(oneRes)-1]
	}
}
