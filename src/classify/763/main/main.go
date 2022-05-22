package main

func partitionLabels(s string) []int {
	// 统计每个字符最后出现的位置
	// 从头遍历字符，并更新字符的最远出现下标，如果找到字符最远出现位置下标和当前下标相等了，则找到了分割点
	mm := make(map[byte]int)
	for i, val := range s {
		mm[byte(val)] = i
	}
	res := make([]int, 0)
	left := 0
	right := 0
	for i, val := range s {
		right = max(right, mm[byte(val)])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
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
