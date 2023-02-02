package main

func commonChars(words []string) []string {
	if len(words) <= 1 {
		return words
	}
	// 先使用一个map记录第一个字符串出现的字符的情况
	firstWordMap := make(map[byte]int)
	for i := 0; i < len(words[0]); i++ {
		firstWordMap[words[0][i]]++
	}
	// 在使用一个map统计其他字符串的字符出现情况，每统计完一个字符串就和第一个字符串的公共元素取最小值
	for i := 1; i < len(words); i++ {
		otherMap := make(map[byte]int)
		for j := 0; j < len(words[i]); j++ {
			otherMap[words[i][j]]++
		}
		//取公共元素最小值
		for k, v := range firstWordMap {
			if otherMap[k] == 0 {
				firstWordMap[k] = 0
			} else {
				firstWordMap[k] = min(otherMap[k], v)
			}
		}
	}
	//统计第一个字符串中剩余的公共元素
	res := make([]string, 0)
	for k, v := range firstWordMap {
		for i := 0; i < v; i++ {
			res = append(res, string(k))
		}
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
