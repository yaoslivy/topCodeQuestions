package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	//逐步判断子串是否可以拼接
	//dp[i] 1-i位置的子串是否可以由wordDict中的词拼接出来
	//dp[i] = true   dp[j] && s[j:i]存在于wordDict中
	//dp[0] = true
	dp := make([]bool, len(s)+1)
	dp[0] = true //空串可以拼接出
	mm := make(map[string]bool)
	for _, val := range wordDict {
		mm[val] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mm[s[j:i]] { // s[j:i] 0 01
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
