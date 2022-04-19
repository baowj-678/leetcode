/**
 * @author: baowj
 * @time: 2022/4/11 09:27:23
 */
package main

func countNumbersWithUniqueDigits(n int) int {
	dp := make([]int, n+1)
	if n == 0 {
		return 0
	}
	dp[0] = 0
	if n == 1 {
		return 10
	}
	dp[1] = 10
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + (dp[i-1]-dp[i-2])*(11-i)
	}
	return dp[n]
}
