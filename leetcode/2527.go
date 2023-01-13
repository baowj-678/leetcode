/**
 * @author: baowj
 * @time: 2023/1/13 11:54:30
 * @email: bwj_678@qq.com
 */
package main

func xorBeauty(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
