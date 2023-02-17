/**
 * @author: baowj
 * @time: 2023/2/16 10:10
 */
package main

import "sort"

func numberOfPairs(nums []int) []int {
	sort.Ints(nums)
	res := []int{0, len(nums)}
	for i := 0; i < len(nums)-1; {
		if nums[i] == nums[i+1] {
			res[0]++
			i += 2
		} else {
			i++
		}
	}
	res[1] -= res[0] * 2
	return res
}
