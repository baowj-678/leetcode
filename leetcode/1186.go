/**
 * @author: baowj
 * @time: 2023/6/27 10:33
 */
package main

import (
	"fmt"
	"math"
)

func maximumSum(arr []int) int {
	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	res := math.MinInt
	for _, x := range arr {
		res = max(res, x)
	}
	if res < 0 {
		return res
	}
	dp := [2]int{0, 0}
	for i, x := range arr {
		if i == 0 {
			dp[0] = x
			dp[1] = 0
		} else {
			dp[0], dp[1] = max(dp[0], 0)+x, max(dp[0], dp[1]+x)
		}
		res = max(res, dp[0])
		res = max(res, dp[1])
	}
	return res
}

func main() {
	arr := []int{1, -4, -5, -2, 5, 0, -1, 2}
	fmt.Println(maximumSum(arr))
}
