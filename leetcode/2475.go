/**
 * @author: baowj
 * @time: 2022/11/22 23:43:04
 * @email: bwj_678@qq.com
 */
package main

func unequalTriplets(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					res++
				}
			}
		}
	}
	return res
}
