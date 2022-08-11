/**
 * @author: baowj
 * @time: 2022/8/9 09:53:00
 */
package main

import "math"

func minStartValue(nums []int) int {
	min := math.MaxInt
	sum := 0
	for num := range nums {
		sum += num
		if min > sum {
			min = sum
		}
	}
	min = 1 - min
	if min < 1 {
		min = 1
	}
	return min
}
