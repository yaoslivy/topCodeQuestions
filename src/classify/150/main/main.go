package main

import "strconv"

func evalRPN(tokens []string) int {
	// 将数字入栈，当遇到算符时，从栈中弹出两个数字，进行运算，然后再压入栈中
	if len(tokens) == 0 {
		return 0
	}
	S := make([]int, 0)
	for _, val := range tokens {
		//遇到算符
		if val == "+" || val == "-" || val == "*" || val == "/" {
			if len(S) < 2 { //错误情况
				return -1
			}
			a := S[len(S)-1]
			S = S[:len(S)-1]
			b := S[len(S)-1]
			S = S[:len(S)-1]
			if val == "+" {
				S = append(S, b+a)
			} else if val == "-" {
				S = append(S, b-a)
			} else if val == "*" {
				S = append(S, b*a)
			} else {
				S = append(S, b/a)
			}
			continue
		}
		//遇到数字
		num, _ := strconv.Atoi(val)
		S = append(S, num)
	}
	return S[len(S)-1]
}
