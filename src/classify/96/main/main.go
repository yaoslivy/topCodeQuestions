package main

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	// dp[i] 表示1-i个节点组成的二叉搜索树种类数
	// dp[i] += dp[j-1] * dp[i-j]  // 以j为根节点，j-1为左子树节点数量，i-j为右子树节点数量
	// dp[0] = 1 ///0个节点，空树也是一种二叉搜索树
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; i-j >= 0; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
