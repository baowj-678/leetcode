/**
 * @author: baowj
 * @time: 2022/12/10 16:57:25
 * @email: bwj_678@qq.com
 */
package main

import "math"

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	dp := make([]bool, 2e4+1)
	for _, x := range baseCosts {
		dp[x] = true
	}
	for _, x := range toppingCosts {
		for i := 2e4 - x; i > 0; i-- {
			if dp[i] {
				dp[i+x] = true
			}
		}

		for i := 2e4 - x; i > 0; i-- {
			if dp[i] {
				dp[i+x] = true
			}
		}
	}
	res := 0
	gap := math.MaxInt
	for i := 1; i < 2e4+1; i++ {
		if dp[i] {
			abs1 := i - target
			if abs1 < 0 {
				abs1 = -abs1
			}
			if abs1 < gap {
				res = i
				gap = abs1
			}
		}
	}
	return res
}

func main() {
	println(closestCost([]int{2, 3}, []int{4, 5, 100}, 18))
}
