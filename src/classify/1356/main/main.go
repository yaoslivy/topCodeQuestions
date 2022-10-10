package main

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		iCnt := bitCount(arr[i])
		jCnt := bitCount(arr[j])
		if iCnt == jCnt {
			return arr[i] < arr[j]
		}
		return iCnt < jCnt
	})
	return arr
}

//位运算计算n的二进制表示中1的个数
func bitCount(n int) int {
	// &= 相同为1， 不同为0
	// 12 : 1100
	//1100 &= 1011 => 1000
	//1000 &= 0111 => 0000
	cnt := 0
	for n != 0 {
		n &= (n - 1)
		cnt++
	}
	return cnt
}
