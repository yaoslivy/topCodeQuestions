package main

func balancedStringSplit(s string) int {
	//最小的平衡字符串即是：包含一个L和一个RR
	//一次遍历，同时统计遍历过的序列中L和R的个数，当两者个数相等时，就分割字符串，然后重新统计L和R的个数
	LCnt, RCnt := 0, 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			RCnt++
		} else {
			LCnt++
		}
		if LCnt != 0 && LCnt == RCnt {
			res++
			LCnt = 0
			RCnt = 0
		}
	}
	return res
}
