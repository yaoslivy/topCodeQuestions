package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	// 局部最优： 遇到strNum[i-1] > strNum[i]的情况，让strNum[i-1]--，然后strNum[i] = 9
	// 全局最优：得到小于等于N的最大单调递增的整数
	strNum := []byte(strconv.Itoa(n))
	flag := len(strNum) //标记赋值9从哪里开始
	for i := len(strNum) - 1; i > 0; i-- {
		if strNum[i-1] > strNum[i] {
			flag = i
			strNum[i-1]--
		}
	}

	for i := flag; i < len(strNum); i++ {
		strNum[i] = '9'
	}
	res, _ := strconv.Atoi(string(strNum))
	return res
}
