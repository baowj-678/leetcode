/**
 * @author: baowj
 * @time: 2022/10/24 22:29:23
 */
package main

func partitionDisjoint(nums []int) int {
	min := make([]int, len(nums))
	min[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < min[i+1] {
			min[i] = nums[i]
		} else {
			min[i] = min[i+1]
		}
	}
	for max, i := nums[0], 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		if max <= nums[i+1] {
			return i
		}
	}
	return 0
}
