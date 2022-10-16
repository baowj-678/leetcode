/**
 * @author: baowj
 * @time: 2022/9/28 08:56:31
 */
package main

func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	dp[0] = 1
	three := 0
	five := 0
	seven := 0
	i := 1
	for i < k {
		if 3*dp[three] <= 5*dp[five] {
			if 3*dp[three] <= 7*dp[seven] {
				dp[i] = 3 * dp[three]
				if 3*dp[three] == 7*dp[seven] {
					seven++
				}
				if 3*dp[three] == 5*dp[five] {
					five++
				}
				three++
			} else {
				dp[i] = 7 * dp[seven]
				seven++
			}
			i++
		} else {
			if 5*dp[five] <= 7*dp[seven] {
				dp[i] = 5 * dp[five]
				if 5*dp[five] == 7*dp[seven] {
					seven++
				}
				five++
			} else {
				dp[i] = 7 * dp[seven]
				seven++
			}
			i++
		}
	}
	return dp[k-1]
}
