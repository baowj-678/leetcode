/**
 * @author: baowj
 * @time: 2023/6/5 15:54
 */
package main

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] == 0 {
			r++
		} else {
			nums[l] = nums[r]
			l++
			r++
		}
	}
	for l < len(nums) {
		nums[l] = 0
		l++
	}
	return nums
}
