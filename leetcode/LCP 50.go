/**
 * @author: baowj
 * @time: 2023/9/15 9:50
 */
package main

import "math"

func giveGem(gem []int, operations [][]int) int {
	minFn := func(x, y int) int {
		if x > y {
			x = y
		}
		return x
	}

	maxFn := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}

	for _, op := range operations {
		gem[op[1]] += gem[op[0]] / 2
		gem[op[0]] -= gem[op[0]] / 2
	}

	min, max := math.MaxInt, 0
	for _, x := range gem {
		min = minFn(min, x)
		max = maxFn(max, x)
	}
	return max - min
}
