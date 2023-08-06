package main

import (
	"math"
	"sort"
)

func deleteGreatestValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}

	res := 0
	for i := 0; i < len(grid[0]); i++ {
		max := math.MinInt
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		res += max
	}
	return res
}
