/**
 * @author: baowj
 * @time: 2022/10/11 23:20:40
 */
package main

func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
	for i := 0; i < len(nums1); i++ {
		nums1[i] -= nums2[i]
	}
	var res int64 = 0
	nums2[len(nums2)-1] = 0
	for i := len(nums1) - 2; i >= 0; i-- {
		num := nums1[i] - diff
		nums2[i] = 0
		for j := i + 1; j < len(nums1); j++ {
			if nums1[j] >= num {
				nums2[i] = nums2[j] + 1
				res += int64(nums2[i])
				break
			}
		}
	}
	return res
}
