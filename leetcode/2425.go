/**
 * @author: baowj
 * @time: 2022/10/11 23:14:17
 */
package main

func xorAllNums(nums1 []int, nums2 []int) int {
	res := 0
	if len(nums1)%2 == 1 {
		for _, num := range nums2 {
			res ^= num
		}
	}
	if len(nums2)%2 == 1 {
		for _, num := range nums1 {
			res ^= num
		}
	}
	return res
}
