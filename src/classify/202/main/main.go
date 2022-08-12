package main

import (
	"math"
	"strconv"
)

func isHappy(n int) bool {
	//使用map记录每个位置上的数字的平方和是否出现，如果下一个出现在map中，则false，如果出现平方和为1则true
	mm := make(map[int]struct{})
	res := n
	for res != 1 {
		res = countN(res)
		if _, ok := mm[res]; ok { //重复出现平方和则false
			return false
		}
		mm[res] = struct{}{}
	}

	return true
}

// 计算数字n每个位置的数字的平方和
func countN(n int) int {
	nLen := len(strconv.Itoa(n))
	res := 0
	for k := 0; k < nLen; k++ {
		temp := n / int(math.Pow(10, float64(k))) % 10 //记录每一位置上的数
		res += temp * temp
	}
	return res
}
