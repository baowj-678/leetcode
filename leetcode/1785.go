/**
 * @author: baowj
 * @time: 2022/12/16 11:05:03
 * @email: bwj_678@qq.com
 */
package main

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	sum -= goal
	if sum < 0 {
		sum = -sum
	}
	return int((int64(sum) + int64(limit-1)) / int64(limit))
}
