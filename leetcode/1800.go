/**
 * @author: baowj
 * @time: 2022/10/8 21:19:46
 */
package main

import "math"

func maxAscendingSum(nums []int) int {
	res := 0
	for i, lastV, cnt := 0, math.MaxInt, 0; i < len(nums); i++ {
		num := nums[i]
		if num > lastV {
			cnt += num
		} else {
			cnt = num
		}
		lastV = num
		if res < cnt {
			res = cnt
		}
	}
	return res
}
