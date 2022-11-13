package main

func repeatedSubstringPattern(s string) bool {
	// return repeatedSubstringPattern2(s)
	return repeatedSubstringPatternKMP(s)
}

func repeatedSubstringPatternKMP(s string) bool {
	if len(s) == 0 {
		return false
	}
	next := getNext(s)
	// 最长相等前后缀abababab
	// 最长相等前缀：ababab
	// 最长相等后缀：  ababab
	// 最长相等前后缀不包含的子串就是最小重复子串
	if next[len(s)-1] != 0 && (len(s)%(len(s)-next[len(s)-1]) == 0) {
		return true
	}
	return false
}

//获取前缀数组
func getNext(s string) []int {
	if len(s) == 0 {
		return []int{}
	}
	j := 0
	next := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}

//暴力求解：枚举所有可能的终点
func repeatedSubstringPatternTraverse(s string) bool {
	sLen := len(s)
	if sLen <= 1 {
		return false
	}
	//所有可能的终点，起点一定是第一个位置
	for i := 0; i < sLen/2; i++ {
		//主串长度需要被子串长度整除
		if sLen%(i+1) != 0 {
			continue
		}
		j := i + 1 //下一个重复子串的开始
		for ; j < sLen; j++ {
			if s[j] != s[j-i-1] { //每次j位置和前i+1个位置比较是否相等
				break
			}
		}
		if j == sLen {
			return true
		}
	}
	return false
}
