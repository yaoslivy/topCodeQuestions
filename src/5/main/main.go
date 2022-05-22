package main

func main() {

}

func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	left := 0
	right := 0
	dp := make([][]bool, len(s)) //dp[i][j]表示s[i...j]是否是回文串
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	//从下到上，从左到右遍历
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] { //i，j位处字符相等，
				if j-i <= 1 { //i，j指向同一个字符或相邻字符则i-j是回文串
					dp[i][j] = true
				} else if dp[i+1][j-1] == true { //中间是否是回文串
					dp[i][j] = true
				}
			}
			if dp[i][j] && j-i+1 > right-left+1 { //判断当前区间i-j的串是否是更大的回文子串
				left = i
				right = j
			}
		}
	}
	return s[left : right+1]
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	//以每一个位置为中心向两边拓展
	for i := 0; i < len(s); i++ {
		left1, right1 := findWidth(s, i, i)   //中心为一个字符的情况
		left2, right2 := findWidth(s, i, i+1) //中心为两个字符的情况
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func findWidth(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1 //注意+1， -1
}
