/**
 * @author: baowj
 * @time: 2023/1/7 21:53:49
 * @email: bwj_678@qq.com
 */
package main

import "math"

func minOperations1658(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	x = sum - x
	if x < 0 {
		return -1
	}
	if x == 0 {
		return len(nums)
	}
	sum = 0
	left, right := 0, 0
	minLen := math.MaxInt
	for ; right < len(nums); right++ {
		sum += nums[right]
		if sum > x {
			for left <= right && sum > x {
				sum -= nums[left]
				left++
			}
		}
		if sum == x {
			if minLen > len(nums)-(right-left+1) {
				minLen = len(nums) - (right - left + 1)
			}
		}
	}
	if minLen == math.MaxInt {
		return -1
	}
	return minLen
}
