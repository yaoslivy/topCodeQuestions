package main

func isIsomorphic(s string, t string) bool {
	// 使用两个map记录映射关系
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; ok && sMap[s[i]] != t[i] {
			return false
		}
		if _, ok := tMap[t[i]]; ok && tMap[t[i]] != s[i] {
			return false
		}
		sMap[s[i]] = t[i]
		tMap[t[i]] = s[i]
	}

	return true
}
