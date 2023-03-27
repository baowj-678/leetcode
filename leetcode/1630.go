/**
 * @author: baowj
 * @time: 2023/3/23 10:00
 */
package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	checkArithmetic := func(l, r int) bool {
		subarray := make([]int, r-l+1)
		copy(subarray, nums[l:r+1])
		sort.Ints(subarray)

		for i := 2; i < len(subarray); i++ {
			if subarray[i]-subarray[i-1] != subarray[i-1]-subarray[i-2] {
				return false
			}
		}
		return true
	}

	res := make([]bool, len(r))
	for i := 0; i < len(res); i++ {
		res[i] = checkArithmetic(l[i], r[i])
	}
	return res
}
