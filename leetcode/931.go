/**
 * @author: baowj
 * @time: 2023/7/13 9:52
 */
package main

import "math"

func minFallingPathSum931(matrix [][]int) int {
	dp := make([]int, len(matrix[0]))
	copy(dp, matrix[0])
	min := func(x, y int) int {
		if x > y {
			x = y
		}
		return x
	}

	for i := 1; i < len(matrix); i++ {
		tmp := make([]int, len(matrix[0]))
		for j := 0; j < len(tmp); j++ {
			tmp[j] = dp[j]
			if j-1 >= 0 {
				tmp[j] = min(tmp[j], dp[j-1])
			}
			if j+1 < len(tmp) {
				tmp[j] = min(tmp[j], dp[j+1])
			}
			tmp[j] += matrix[i][j]
		}
		dp = tmp
	}
	res := math.MaxInt
	for _, x := range dp {
		res = min(res, x)
	}
	return res
}
