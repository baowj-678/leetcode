/**
 * @author: baowj
 * @time: 2022/9/17 20:29:53
 */
package main

const MOD2400 = 1e9 + 7

func numberOfWays(startPos int, endPos int, k int) int {
	dp1 := make([]int, 2005)
	dp2 := make([]int, 2005)
	start := k + 2
	dp1[start] = 1
	for i := 0; i < k; i++ {
		for j := 1; j < len(dp1)-1; j++ {
			dp2[j] = (dp1[j-1] + dp1[j+1]) % MOD2400
		}
		dp1, dp2 = dp2, dp1
	}
	return dp1[start+endPos-startPos]
}
