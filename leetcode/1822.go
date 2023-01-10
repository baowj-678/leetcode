/**
 * @author: baowj
 * @time: 2022/10/27 16:57:50
 */
package main

func arraySign(nums []int) int {
	res := 1
	for _, num := range nums {
		if num < 0 {
			res = -res
		} else if num == 0 {
			return 0
		}
	}
	return res
}
