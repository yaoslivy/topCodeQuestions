package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	// 局部最优： 遇到strNum[i-1] > strNum[i]的情况，让strNum[i-1]--，然后strNum[i] = 9
	// 全局最优：得到小于等于N的最大单调递增的整数
	// 从后往前，让前一个比后一个大的位置上的值--，保证了改变后的值比原来值小
	// 记录第一个--的位置，让后一个位置开始都赋值为9，这样保证了值尽可能的大
	strNum := []byte(strconv.Itoa(n))
	flag := len(strNum) //记录第一个--的位置
	for i := len(strNum) - 1; i > 0; i-- {
		if strNum[i-1] > strNum[i] {
			flag = i - 1
			strNum[i-1]--
		}
	}

	for i := flag + 1; i < len(strNum); i++ {
		strNum[i] = '9'
	}
	res, _ := strconv.Atoi(string(strNum))
	return res
}
