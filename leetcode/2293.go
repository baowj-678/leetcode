/**
 * @author: baowj
 * @time: 2023/1/23 17:04:08
 * @email: bwj_678@qq.com
 */
package main

func minMaxGame(nums []int) int {
	for len(nums) > 1 {
		newNums := make([]int, len(nums)/2)
		for i := 0; i < len(newNums); i++ {
			if i%2 == 0 {
				if nums[i*2] > nums[2*i+1] {
					newNums[i] = nums[2*i+1]
				} else {
					newNums[i] = nums[2*i]
				}
			} else {
				if nums[i*2] > nums[2*i+1] {
					newNums[i] = nums[2*i]
				} else {
					newNums[i] = nums[2*i+1]
				}
			}
		}
		nums = newNums
	}
	return nums[0]
}
