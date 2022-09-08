package main

import "strings"

func reverseWords(s string) string {
	if len(s) == 0 {
		return " "
	}
	//去除两边空格
	s = strings.Trim(s, " ")
	//中间重复空格保留一个
	for i := 0; i < len(s)-1; {
		if s[i] == ' ' && s[i+1] == ' ' {
			s = s[:i] + s[i+1:]
		} else {
			i++
		}
	}
	sArr := strings.Split(s, " ")
	reverseStrArr(sArr, 0, len(sArr)-1)

	return strings.Join(sArr, " ")
}

func reverseStrArr(sArr []string, left, right int) {
	for left < right {
		sArr[left], sArr[right] = sArr[right], sArr[left]
		left++
		right--
	}
}

func reverseWords2(s string) string {
	//先去除前后两边的空格
	for len(s) != 0 && s[0] == ' ' {
		s = s[1:]
	}
	for len(s) != 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	//去除中间空格
	i := 0
	for i < len(s) {
		if i > 0 && i < len(s)-1 && s[i] == ' ' && s[i+1] == ' ' {
			s = s[:i] + s[i+1:]
		} else {
			i++
		}
	}
	sArr := []byte(s)
	//对剩下的字符先做一个全局反转，再逐个对单个单词做反转
	reverse(sArr, 0, len(sArr)-1)
	i, j := 0, 0
	for j < len(sArr) {
		for j < len(sArr) && sArr[j] != ' ' {
			j++
		}
		if j >= len(sArr) {
			break
		}
		reverse(sArr, i, j-1)
		i = j + 1
		j++
	}
	//将最后一个单词反转
	reverse(sArr, i, j-1)
	return string(sArr)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
