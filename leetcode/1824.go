/**
 * @author: baowj
 * @time: 2023/1/21 16:10:56
 * @email: bwj_678@qq.com
 */
package main

import (
	"fmt"
	"math"
)

func minSideJumps(obstacles []int) int {
	dp := make([][3]int, len(obstacles))
	dp[0][0], dp[0][1], dp[0][2] = 1, 0, 1
	for i := 1; i < len(obstacles); i++ {
		dp[i] = dp[i-1]
		if obstacles[i] != 0 {
			dp[i][obstacles[i]-1] = math.MaxInt
		}
		min := math.MaxInt
		for _, x := range dp[i] {
			if min > x {
				min = x
			}
		}
		for j := 0; j < 3; j++ {
			if dp[i][j] > min+1 {
				dp[i][j] = min + 1
			}
		}
		if obstacles[i] != 0 {
			dp[i][obstacles[i]-1] = math.MaxInt
		}
	}
	fmt.Println(dp)
	min := math.MaxInt
	for _, x := range dp[len(dp)-1] {
		if min > x {
			min = x
		}
	}
	return min
}
