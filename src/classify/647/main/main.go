package main

func countSubstrings(s string) int {
	// dp[i][j] 表示区间[i,j]内的子串是否是回文子串 true/false
	// dp[i][j] = true if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) //两指针指向的字符相等且区间内是回文子串
	// dp[][] = false
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}
	}
	return res
}

func countSubstrings2(s string) int {
	res := 0
	for i := 0; i < len(s); i++ { //以各个位置向左右两边扩，观察是否为回文子串
		res += find(s, i, i)   // 以i为中心
		res += find(s, i, i+1) // 以i和i+1为中心
	}
	return res
}

func find(s string, i, j int) int {
	res := 0
	for i >= 0 && j <= len(s)-1 {
		if s[i] == s[j] {
			res++
			i--
			j++
		} else {
			break
		}
	}
	return res
}
