package main

func partitionLabels(s string) []int {
	// 统计每个字符最后出现的位置
	// 从头遍历字符，并更新字符的最远出现下标，
	// 如果找到字符最远出现位置下标和当前下标相等了，则找到了分割点
	sArr := []byte(s)
	mm := make(map[byte]int)
	for i, val := range sArr {
		mm[val] = i
	}
	//记录当前区间的左右边界
	left := 0
	right := 0
	res := make([]int, 0)
	// ababcbaca defegdehijhklij
	// a : 8
	for i := 0; i < len(sArr); i++ {
		right = max(right, mm[sArr[i]]) //当前最远距离
		if i >= right {                 //到达当前区间所有字符的最远距离
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
