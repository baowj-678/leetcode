/**
 * @author: baowj
 * @time: 2022/12/10 21:39:05
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func minOperations1775(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	sum1, sum2 := 0, 0
	for _, x := range nums1 {
		sum1 += x
	}
	for _, x := range nums2 {
		sum2 += x
	}
	if sum1 < sum2 {
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}
	gap := sum1 - sum2
	p1, p2 := len(nums1)-1, 0
	res := 0
	for p1 >= 0 && p2 < len(nums2) && gap > 0 {
		if nums1[p1]-1 > 6-nums2[p2] {
			if nums1[p1]-1 >= gap {
				return res + 1
			} else {
				gap -= nums1[p1] - 1
				res++
			}
			p1--
		} else {
			if 6-nums2[p2] >= gap {
				return res + 1
			} else {
				gap -= 6 - nums2[p2]
				res++
			}
			p2++
		}
	}
	if gap > 0 {
		for p1 >= 0 {
			gap -= nums1[p1] - 1
			res++
			if gap <= 0 {
				return res
			}
			p1--
		}

		for p2 < len(nums2) {
			gap -= 6 - nums2[p2]
			res++
			if gap <= 0 {
				return res
			}
			p2++
		}
	} else {
		return res
	}
	return -1
}
