/**
 * @author: baowj
 * @time: 2023/9/5 10:34
 */
package main

import (
	"math"
	"sort"
)

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	v1, v2 := nums1[0], nums2[0]
	i, j := 0, 0
	res := math.MaxInt
	min := func(x, y int) int {
		if x > y {
			x = y
		}
		return x
	}
	for i < len(nums1) && j < len(nums2) {
		v1 = nums1[i]
		v2 = nums2[j]
		res = min(res, min(v1*10+v2, v2*10+v1))
		if v1 == v2 {
			return v1
		}
		if v1 < v2 {
			i++
		} else {
			j++
		}
	}
	return res
}
