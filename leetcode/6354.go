/**
 * @author: baowj
 * @time: 2023/2/12 16:34:37
 * @email: bwj_678@qq.com
 */
package main

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	var res int64 = 0
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		val := strconv.Itoa(nums[i])
		if i != j {
			val += strconv.Itoa(nums[j])
		}
		x, _ := strconv.Atoi(val)
		res += int64(x)
	}
	return res
}
