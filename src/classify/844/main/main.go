package main

func backspaceCompare(s string, t string) bool {
	//双指针思路
	i := len(s) - 1
	j := len(t) - 1
	sCnt, tCnt := 0, 0
	for {
		// 从后往前，消除s的#
		for i >= 0 {
			if s[i] == '#' {
				sCnt++
			} else {
				if sCnt > 0 {
					sCnt--
				} else {
					break
				}
			}
			i--
		}
		//从后往前，消除t的#
		for j >= 0 {
			if t[j] == '#' {
				tCnt++
			} else {
				if tCnt > 0 {
					tCnt--
				} else {
					break
				}
			}
			j--
		}
		//当前i,j位置之后的#都消除完毕
		if i < 0 || j < 0 {
			break
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	if i == -1 && j == -1 {
		return true
	}
	return false
}

func backspaceCompare2(s string, t string) bool {
	// 双指针思路
	// 都从后往前，遇到#统计个数，消除前面的字符，
	// 没有#时，则比较字符是否相等，不等则false
	i, j := len(s)-1, len(t)-1
	sCnt, tCnt := 0, 0
	for i >= 0 && j >= 0 {
		//统计#的同时消除#
		for i >= 0 {
			if s[i] == '#' {
				sCnt++
			} else {
				if sCnt > 0 {
					sCnt--
				} else {
					break
				}
			}
			i--
		}
		for j >= 0 {
			if t[j] == '#' {
				tCnt++
			} else {
				if tCnt > 0 {
					tCnt--
				} else {
					break
				}
			}
			j--
		}
		if i >= 0 && j >= 0 && s[i] == t[j] {
			i--
			j--
		} else {
			break
		}
	}
	// 如果i或者j前面还有#，则统计，并清除
	//统计#的同时消除#
	for i >= 0 {
		if s[i] == '#' {
			sCnt++
		} else {
			if sCnt > 0 {
				sCnt--
			} else {
				break
			}
		}
		i--
	}
	for j >= 0 {
		if t[j] == '#' {
			tCnt++
		} else {
			if tCnt > 0 {
				tCnt--
			} else {
				break
			}
		}
		j--
	}
	if i == -1 && j == -1 {
		return true
	}
	return false
}
