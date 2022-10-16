/**
 * @author: baowj
 * @time: 2022/9/25 20:54:50
 */
package main

import (
	"math"
)

func goodIndices(nums []int, k int) []int {
	if k == 1 {
		res := make([]int, len(nums)-2)
		for i := 0; i < len(nums)-2; i++ {
			res[i] = i + 1
		}
		return res
	}
	noIncreaseFlag := make([]bool, len(nums))
	noDecreaseFlag := make([]bool, len(nums))
	lastVal := math.MaxInt
	length := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= lastVal {
			length++
		} else {
			length = 1
		}
		lastVal = nums[i]
		if length >= k {
			noIncreaseFlag[i] = true
		}
	}
	lastVal = math.MaxInt
	length = 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= lastVal {
			length++
		} else {
			length = 1
		}
		lastVal = nums[i]
		if length >= k {
			noDecreaseFlag[i] = true
		}
	}
	res := []int{}
	for i := k; i < len(nums)-k; i++ {
		if noIncreaseFlag[i-1] && noDecreaseFlag[i+1] {
			res = append(res, i)
		}
	}
	return res
}
