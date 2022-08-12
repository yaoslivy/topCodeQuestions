package main

func reverseWords(s string) string {
	if len(s) <= 1 {
		return s
	}
	i := 0
	//去除前后空格
	for i < len(s) && s[i] == ' ' {
		s = s[i+1:]
	}
	j := len(s) - 1
	for j >= 0 && s[j] == ' ' {
		s = s[:j]
		j--
	}
	i = 0
	for i < len(s) {
		for i < len(s) && s[i] == ' ' && s[i+1] == ' ' {
			s = s[:i] + s[i+1:] //去除i位置的空格
		}
		i++
	}

	// 将整个字符串反转，再逐步将每个单词反转
	strs := []byte(s)
	reverse(strs, 0, len(strs)-1)
	//world hello
	i, j = 0, 0
	for j < len(strs) {
		for j < len(strs) && strs[j] != ' ' {
			j++
		}
		if j >= len(strs) {
			break
		}
		reverse(strs, i, j-1)
		i = j + 1
		j++
	}
	reverse(strs, i, j-1)
	return string(strs)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
