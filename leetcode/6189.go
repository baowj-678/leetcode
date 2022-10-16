/**
 * @author: baowj
 * @time: 2022/9/25 20:39:39
 */
package main

func longestSubarray(nums []int) int {
	max := 0
	for _, x := range nums {
		if x > max {
			max = x
		}
	}
	res := 0
	tmp := 0
	for _, x := range nums {
		if x == max {
			tmp++
		} else {
			tmp = 0
		}
		if res < tmp {
			res = tmp
		}
	}
	return res
}
