package main

func strStr(haystack string, needle string) int {
	// return strStrTraverse(haystack, needle)
	return strStrKMP(haystack, needle)
}

// KMP算法
func strStrKMP(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := getNext(needle)
	j := 0 //子串起始匹配位置
	for i := 0; i < len(haystack); i++ {
		//当前字符不匹配，则找前一个位置对应的回退位置
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

//构造next数组 获得前缀表，前缀表用来记录模式串与主串不匹配时，模式串应该从哪里开始重新匹配
// 字符串的前缀是指：不包含最后一个字符的所有以第一个字符开头的连续子串
// 字符串的后缀是指：不包含第一个字符的所有以最后一个字符结尾的连续子串
//获取模式串的next数组  aabaa 0 1
func getNext(s string) []int {
	next := make([]int, len(s))
	j := 0 //前缀相等地方
	//找到每一个位置的最长公共前后缀长度，第一位为0
	// aabaaf // 0 1 0 1 2 0
	for i := 1; i < len(s); i++ { //需要填充的位置
		//如果当前j位置和i位置不等，则找到j-1前一个位置对应的回退位置
		for j > 0 && s[j] != s[i] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
	return next
}

// 暴力遍历
func strStrTraverse(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	//遍历所有haystack字符为起点，查看是否和needle匹配
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		//当前起点
		start := i
		j := 0
		for start < len(haystack) && j < len(needle) && haystack[start] == needle[j] {
			start++
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
