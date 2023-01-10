package main

const MAX_790 = 1e9 + 7

func numTilings(n int) int {
	dp := make([][3]int, n+1)
	dp[0] = [3]int{0, 0, 1}
	dp[1] = [3]int{0, 0, 1}
	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-2][2] + dp[i-1][1]) % MAX_790
		dp[i][1] = (dp[i-2][2] + dp[i-1][0]) % MAX_790
		dp[i][2] = (dp[i-1][2] + dp[i-1][0] + dp[i-1][1] + dp[i-2][2]) % MAX_790
	}
	return dp[n][2]
}
