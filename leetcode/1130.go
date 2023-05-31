/**
 * @author: baowj
 * @time: 2023/5/31 9:53
 */
package main

import "math"

func mctFromLeafValues(arr []int) int {
	m := make([][]int, len(arr))
	// init max matrix
	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	for i := 0; i < len(arr); i++ {
		m[i] = make([]int, len(arr))
		maxV := arr[i]
		for j := i; j < len(arr); j++ {
			maxV = max(maxV, arr[j])
			m[i][j] = maxV
		}
	}
	// dp
	min := func(x, y int) int {
		if x > y {
			x = y
		}
		return x
	}
	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, len(arr))
	}
	for l := 1; l < len(arr); l++ {
		for i := 0; i+l < len(arr); i++ {
			j := i + l
			minV := math.MaxInt
			for mid := i; mid < j; mid++ {
				tmp := dp[i][mid] + dp[mid+1][j] + m[i][mid]*m[mid+1][j]
				minV = min(tmp, minV)
			}
			dp[i][j] = minV
		}
	}
	return dp[0][len(arr)]
}
