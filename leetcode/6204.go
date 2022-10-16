/**
 * @author: baowj
 * @time: 2022/10/16 18:57:12
 */
package main

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		if -nums[l] == nums[r] {
			return nums[r]
		} else {
			if -nums[l] < nums[r] {
				r--
			} else {
				l++
			}
		}
	}
	return -1
}
