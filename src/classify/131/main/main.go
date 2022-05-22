package main

import "fmt"

func main() {
	s := "aab"
	ans := partition(s)
	for _, val := range ans {
		fmt.Println(val)
	}
}

var res [][]string

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	res = make([][]string, 0)
	find(s, 0, []string{})
	return res
}

//将切割问题变为组合问题，
func find(s string, start int, oneRes []string) {
	// 截取字符串的起始位置已经在最后一个元素之后了，说明path中的子串都是回文串，并且是s的所有子串
	if start == len(s) {
		copyArr := make([]string, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}

	for i := start; i < len(s); i++ { //不断的尝试截取start:i+1
		if valid(s[start : i+1]) {
			oneRes = append(oneRes, s[start:i+1])
			find(s, i+1, oneRes)
			oneRes = oneRes[:len(oneRes)-1] //去除本次填充的子串
		}
	}
}

//判断s是否是回文串
func valid(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
