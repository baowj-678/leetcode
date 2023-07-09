/**
 * @author: baowj
 * @time: 2023/6/16 14:15
 */
package main

func wiggleMaxLength(nums []int) int {
	up := make([]int, len(nums))
	down := make([]int, len(nums))
	up[0], down[0] = 1, 1
	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			up[i] = up[i-1]
			down[i] = down[i-1]
		} else if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else {
			down[i] = max(down[i-1], up[i-1]+1)
			up[i] = up[i-1]
		}
	}
	return max(up[len(nums)-1], down[len(nums)-1])
}
