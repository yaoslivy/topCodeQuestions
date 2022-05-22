package main

func fib2(n int) int {
	dp := make([]int, 2) //dp[i]表示第i个数的斐波那契值
	if n == 0 || n == 1 {
		return n
	}

	//初始化
	dp[0] = 0
	dp[1] = 1
	var sum int
	for i := 2; i <= n; i++ { //从左到右依次填充
		sum = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}

	return dp[1]
}

func fib(n int) int {
	dp := make([]int, n+1) //dp[i]表示第i个数的斐波那契值
	if n == 0 || n == 1 {
		return n
	}

	//初始化
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ { //从左到右依次填充
		dp[i] = dp[i-1] + dp[i-2] //递推公式
	}

	return dp[n]
}
