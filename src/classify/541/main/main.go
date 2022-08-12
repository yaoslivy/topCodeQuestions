package main

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
