package main

func isValid2(s string) bool {
	// 长度必须可以是2的倍数
	if len(s)%2 != 0 {
		return false
	}
	// 遍历序列，遇到左括号入栈，遇到右括号，出栈
	S := make([]rune, 0)
	for _, val := range s {
		//左括号情况
		if val == '(' || val == '{' || val == '[' {
			S = append(S, val)
			continue
		}
		//右括号情况，栈中无元素或栈顶元素不等于右括号匹配的左括号则返回false
		if len(S) == 0 {
			return false
		}
		if S[len(S)-1] == '(' && val == ')' {
			S = S[:len(S)-1]
			continue
		}
		if S[len(S)-1] == '{' && val == '}' {
			S = S[:len(S)-1]
			continue
		}
		if S[len(S)-1] == '[' && val == ']' {
			S = S[:len(S)-1]
			continue
		}
		return false
	}
	// 如果有多余的左括号
	if len(S) != 0 {
		return false
	}
	return true
}

func isValid(s string) bool {
	// 使用一个栈记录遍历过的左字符串
	S := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			S = append(S, s[i])
		} else {
			if s[i] == ')' && (len(S) == 0 || S[len(S)-1] != '(') {
				return false
			}
			if s[i] == ']' && (len(S) == 0 || S[len(S)-1] != '[') {
				return false
			}
			if s[i] == '}' && (len(S) == 0 || S[len(S)-1] != '{') {
				return false
			}
			S = S[:len(S)-1]
		}
	}
	if len(S) != 0 {
		return false
	}
	return true
}
