/**
 * @author: baowj
 * @time: 2022/11/24 21:59:57
 * @email: bwj_678@qq.com
 */
package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	lastLarge, lastEqual := -1, -1
	res := 0
	for i, num := range nums {
		if num > right {
			lastLarge = i
			lastEqual = -1
			continue
		} else if num < left {
			if lastEqual != -1 {
				res += lastEqual - lastLarge
			}
		} else {
			lastEqual = i
			res += lastEqual - lastLarge
		}
	}
	return res
}
