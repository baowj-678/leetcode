package main

import "math"

func minFallingPathSum(grid [][]int) int {
	dp := grid[0]
	min := func(x, y int) int {
		if x > y {
			x = y
		}
		return x
	}

	for i := 1; i < len(grid); i++ {
		tmp := make([]int, len(grid[i]))
		for j := 0; j < len(tmp); j++ {
			tmp[j] = math.MaxInt
			for k := 0; k < len(grid[i]); k++ {
				if k == j {
					continue
				}
				tmp[j] = min(tmp[j], dp[k]+grid[i][j])
			}
		}
		dp = tmp
	}

	res := math.MaxInt
	for _, x := range dp {
		res = min(res, x)
	}
	return res
}
