package main

func removeDuplicates(s string) string {
	//将元素入栈，如何当前元素和栈顶元素相同则将当前元素和栈顶元素一并删除，下一个元素也是如此
	if len(s) <= 1 {
		return s
	}
	S := ""
	for _, val := range s {
		//栈中元素为空则入栈
		if len(S) == 0 {
			S += string(val)
			continue
		}
		//当前元素和栈顶元素相同，则一同出栈
		if byte(val) == S[len(S)-1] {
			S = S[:len(S)-1]
			continue
		}
		//当前元素和栈顶元素不同
		S += string(val)
	}
	return S
}
