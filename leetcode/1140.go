/**
 * @author: baowj
 * @time: 2023/2/26 15:54:53
 * @email: bwj_678@qq.com
 */
package main

import "math"

func stoneGameII(piles []int) int {
	dp := make([][]int, len(piles))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(piles))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	sum := 0
	for _, pile := range piles {
		sum += pile
	}
	return (sum + dfs1140(piles, dp, 0, 1)) / 2
}

func dfs1140(piles []int, dp [][]int, i, m int) int {
	if dp[i][m] != -1 {
		return dp[i][m]
	}
	max := math.MinInt
	sum := 0
	if i+2*m >= len(piles) {
		for k := i; k < len(piles); k++ {
			sum += piles[k]
		}
		return sum
	}
	for k := 1; k <= 2*m && k+i < len(piles); k++ {
		sum += piles[i+k-1]
		newM := k
		if newM < m {
			newM = m
		}
		if sum-dfs1140(piles, dp, i+k, newM) > max {
			max = sum - dfs1140(piles, dp, i+k, newM)
		}
	}
	dp[i][m] = max
	return max
}
