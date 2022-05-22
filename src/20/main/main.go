package main

func isValid(s string) bool {
	S := make([]byte, 0)
	for _, val := range s {
		switch val {
		case '(':
			S = append(S, byte(val))
		case '{':
			S = append(S, byte(val))
		case '[':
			S = append(S, byte(val))
		case ')':
			if len(S) == 0 || S[len(S)-1] != '(' {
				return false
			}
			S = S[:len(S)-1]
		case '}':
			if len(S) == 0 || S[len(S)-1] != '{' { //注意栈中元素为0的情况
				return false
			}
			S = S[:len(S)-1]
		case ']':
			if len(S) == 0 || S[len(S)-1] != '[' {
				return false
			}
			S = S[:len(S)-1]
		}
	}
	if len(S) != 0 { //注意栈中还存在元素的情况
		return false
	}
	return true
}
