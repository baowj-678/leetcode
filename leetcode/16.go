/**
 * @author: baowj
 * @time: 2023/7/10 10:58
 */
package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := int(1e9)
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}

	for i := 0; i < len(nums); i++ {
		l, r := 0, len(nums)-1

		for l < r {
			if l == i {
				l++
				continue
			}
			if r == i {
				r--
				continue
			}

			sum := nums[i] + nums[l] + nums[r]

			if abs(sum-target) < abs(res-target) {
				res = sum
			}

			if res == target {
				return res
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
