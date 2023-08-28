/**
 * @author: baowj
 * @time: 2023/7/20 10:06
 */
package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	res := math.MinInt

	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}

	sum := 0
	leftSum := 0
	leftMax := math.MinInt
	leftPrefix := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res = max(res, sum)

		leftSum += nums[i]
		leftMax = max(leftMax, leftSum)
		leftPrefix[i] = leftMax
		if sum < 0 {
			sum = 0
		}
	}

	sum = 0
	rightMax := math.MinInt
	for i := len(nums) - 1; i > 0; i-- {
		sum += nums[i]
		rightMax = max(rightMax, sum)
		res = max(res, rightMax+leftPrefix[i-1])
	}
	return res
}
