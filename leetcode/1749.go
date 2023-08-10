package main

import "math"

func maxAbsoluteSum(nums []int) int {
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}

	maxF := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}

	minF := func(x, y int) int {
		if x > y {
			x = y
		}
		return x
	}
	max, min := math.MinInt, math.MaxInt
	maxSum, minSum := 0, 0
	for _, num := range nums {
		maxSum += num
		minSum += num
		if maxSum < 0 {
			maxSum = 0
		}
		if minSum > 0 {
			minSum = 0
		}
		max = maxF(max, maxSum)
		min = minF(min, minSum)
	}
	return maxF(abs(max), abs(min))
}
