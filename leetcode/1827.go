/**
 * @author: baowj
 * @time: 2022/12/16 11:09:29
 * @email: bwj_678@qq.com
 */
package main

func minOperations1827(nums []int) int {
	res := 0
	last := 0
	for _, n := range nums {
		if last >= n {
			res += last + 1 - n
			n = last + 1
		}
		last = n
	}
	return res
}
