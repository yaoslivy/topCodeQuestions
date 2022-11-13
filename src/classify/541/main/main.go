package main

func reverseStr2(s string, k int) string {
	if k <= 0 {
		return s
	}
	sArr := []byte(s)
	//需要反转的右边界
	//ab cd efg
	for i := k - 1; i < len(s)+k; i += 2 * k {
		if i < len(s) {
			reverse(sArr, i-k+1, i)
		} else {
			reverse(sArr, i-k+1, len(s)-1)
		}
	}
	return string(sArr)
}

func reverseStr(s string, k int) string {
	if len(s) == 0 {
		return s
	}
	resBytes := []byte(s)
	//abcd efg
	for i := 0; i < len(resBytes); i = i + 2*k {
		if i+k < len(s) {
			reverse(resBytes, i, i+k-1)
		} else {
			reverse(resBytes, i, len(s)-1)
		}
	}
	return string(resBytes)
}

func reverse(s []byte, left int, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
