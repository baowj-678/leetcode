/**
 * @author: baowj
 * @time: 2022/11/18 17:12
 */
package main

var MOD2466 int = 1e9 + 7

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	res := 0
	for i := 1; i <= high; i++ {
		if i-zero >= 0 {
			dp[i] = (dp[i] + dp[i-zero]) % MOD2466
		}
		if i-one >= 0 {
			dp[i] = (dp[i] + dp[i-one]) % MOD2466
		}
		if i >= low {
			res = (res + dp[i]) % MOD2466
		}
	}
	return res
}
