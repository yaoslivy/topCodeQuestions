package main

func longestPalindrome(s string) string {
	//动态规划
	// return longestPalindromeDynamic(s)
	//双指针方法
	return longestPalindromeDoublePointer(s)
}

func longestPalindromeDoublePointer(s string) string {
	if len(s) <= 1 {
		return s
	}
	start, end := 0, 0
	//以每个位置为可能的中心，向两边拓展回文子串
	for i := 0; i < len(s); i++ {
		left1, right1 := getLongestLen(s, i, i)
		left2, right2 := getLongestLen(s, i, i+1)
		if right2-left2 > right1-left1 {
			left1, right1 = left2, right2
		}
		if right1-left1 > end-start {
			start, end = left1, right1
		}
	}
	return s[start : end+1]
}

func getLongestLen(s string, i, j int) (int, int) {
	start, end := 0, 0
	for i >= 0 && j < len(s) {
		if s[i] == s[j] {
			if end-start < j-i {
				start, end = i, j
			}
			i--
			j++
		} else {
			break
		}
	}
	return start, end
}

func longestPalindromeDynamic(s string) string {
	if len(s) <= 1 {
		return s
	}
	//dp[i][j] 表示i到j的子串是否是回文串
	//dp[i][j] = true if s[i] == s[j] && (dp[i+1][j-1] || i+1==j)
	// dp[i][i] = true
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	start, end := 0, 0
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && (dp[i+1][j-1] || i+1 == j) {
				dp[i][j] = true
				if j-i+1 > end-start+1 {
					end, start = j, i
				}
			}
		}
	}

	return s[start : end+1]
}
