package main

var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = make([]string, 0)
	digitsStr := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	find(digits, 0, "", digitsStr)
	return res
}

func find(digits string, start int, oneRes string, digitsStr map[string]string) {
	if len(oneRes) == len(digits) { //当选取到的字符长度符合条件时
		res = append(res, oneRes)
		return
	}

	for i := start; i < len(digits); i++ { //每次选取一个位置
		str, _ := digitsStr[string(digits[i])]
		for j := 0; j < len(str); j++ { //尝试选取一个字符
			oneRes = oneRes + string(str[j])
			find(digits, i+1, oneRes, digitsStr) //每次在当前数字对应的字符序列中选取一个字符
			oneRes = oneRes[:len(oneRes)-1]
		}
	}
}
