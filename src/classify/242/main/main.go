package main

func isAnagram(s string, t string) bool {
	//使用map记录s中所有字符的出现个数，再遍历t的过程中，如果t中的字符没有在map中则false
	if len(s) != len(t) {
		return false
	}
	mm := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mm[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if mm[t[i]] == 0 {
			return false
		}
		mm[t[i]]--
	}
	return true
}
