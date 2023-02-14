/**
 * @author: baowj
 * @time: 2022/11/18 17:09
 */
package main

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	set := map[int]struct{}{}
	for len(nums) > 0 {
		sum := nums[0] + nums[len(nums)-1]
		set[sum] = struct{}{}
		nums = nums[1 : len(nums)-1]
	}
	return len(set)
}
