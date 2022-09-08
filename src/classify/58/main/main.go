package main

func reverseLeftWords(s string, n int) string {
	// 先局部反转0:n，n: ；再整体反转
	if len(s) < n {
		return s
	}
	strs := []byte(s)
	reverse(strs[:n])
	reverse(strs[n:])
	reverse(strs)
	return string(strs)
}

func reverse(s []byte) {
	if len(s) <= 1 {
		return
	}
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseLeftWords2(s string, n int) string {
	if len(s) < n {
		return s
	}
	return s[n:] + s[:n]
}
