package main

func canConstruct(ransomNote string, magazine string) bool {
	// 使用长度为26数组记录每个字符出现的个数
	res := make([]rune, 26)
	for _, val := range magazine {
		res[val-'a']++
	}
	for _, val := range ransomNote {
		if res[val-'a'] < 1 {
			return false
		}
		res[val-'a']--
	}
	return true
}

// 直接使用map
func canConstruct2(ransomNote string, magazine string) bool {
	// 遍历一次magazine，map记录各个字符出现的个数，在遍历ransomNote的同时判断是否可以从magazine中取出字符构成
	if len(ransomNote) > len(magazine) {
		return false
	}
	mm := make(map[rune]int)
	for _, val := range magazine {
		mm[val]++
	}

	for _, val := range ransomNote {
		if mm[val] < 1 {
			return false
		}
		mm[val]--
	}
	return true
}
