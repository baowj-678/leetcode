/**
 * @author: baowj
 * @time: 2022/8/29 09:21:34
 */
package main

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[i*2] = nums[i]
		res[i*2+1] = nums[n+i]
	}
	return res
}
