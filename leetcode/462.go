package main

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	v := nums[len(nums)/2]
	res := 0
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for _, num := range nums {
		res += abs(num - v)
	}
	return res
}
