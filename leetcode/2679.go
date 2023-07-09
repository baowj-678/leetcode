/**
 * @author: baowj
 * @time: 2023/7/9 14:47
 */
package main

import "sort"

func matrixSum(nums [][]int) int {
	findMax := func(idx int) int {
		res := nums[0][idx]
		for i := 1; i < len(nums); i++ {
			if res < nums[i][idx] {
				res = nums[i][idx]
			}
		}
		return res
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
	}
	for i := 0; i < len(nums[0]); i++ {
		res += findMax(i)
	}
	return res
}
