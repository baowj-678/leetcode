/**
 * @author: baowj
 * @time: 2022/10/9 16:40:28
 */
package main

const MOD6203 = 1e9 + 7

func numberOfPaths(grid [][]int, k int) int {
	dp := make([][][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, len(grid[0]))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, k)
		}
	}
	dp[0][0][grid[0][0]%k] = 1
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i > 0 {
				for t := 0; t < k; t++ {
					idx := (grid[i][j] + t) % k
					dp[i][j][idx] = (dp[i][j][idx] + dp[i-1][j][t]) % MOD6203
				}
			}
			if j > 0 {
				for t := 0; t < k; t++ {
					idx := (grid[i][j] + t) % k
					dp[i][j][idx] = (dp[i][j][idx] + dp[i][j-1][t]) % MOD6203
				}
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1][0]
}
