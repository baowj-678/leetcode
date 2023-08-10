package main

import "sort"

func maxHeight(cuboids [][]int) int {
	for i := 0; i < len(cuboids); i++ {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] == cuboids[j][0] {
			if cuboids[i][1] == cuboids[j][1] {
				return cuboids[i][2] < cuboids[j][2]
			} else {
				return cuboids[i][1] < cuboids[j][1]
			}
		} else {
			return cuboids[i][0] < cuboids[j][0]
		}
	})

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	res := 0
	dp := make([]int, len(cuboids))
	for i := 0; i < len(cuboids); i++ {
		maxH := cuboids[i][2]
		for j := 0; j < i; j++ {
			if cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
				maxH = max(maxH, dp[j]+cuboids[i][2])
			}
		}
		dp[i] = maxH
		res = max(res, maxH)
	}
	return res
}
