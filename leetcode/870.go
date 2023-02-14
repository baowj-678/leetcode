package main

import (
	"sort"
)

type struct870 struct {
	nums []int
	idxs []int
}

func (s struct870) Len() int           { return len(s.nums) }
func (s struct870) Less(i, j int) bool { return s.nums[i] <= s.nums[j] }
func (s struct870) Swap(i, j int) {
	s.nums[i], s.nums[j], s.idxs[i], s.idxs[j] = s.nums[j], s.nums[i], s.idxs[j], s.idxs[i]
}

func advantageCount(nums1 []int, nums2 []int) []int {
	idxs := make([]int, len(nums2))
	for i := 0; i < len(idxs); i++ {
		idxs[i] = i
	}
	s := struct870{
		nums: nums2,
		idxs: idxs,
	}
	sort.Sort(s)
	sort.Ints(nums1)
	res := make([]int, len(nums1))
	for i, j, tail := 0, 0, len(nums1)-1; j <= tail && i < len(nums1) && j < len(nums2); {
		if nums1[i] > nums2[j] {
			res[j] = nums1[i]
			i++
			j++
		} else {
			res[tail] = nums1[i]
			i++
			tail--
		}
	}
	for i := 0; i < len(nums1); i++ {
		nums1[idxs[i]] = res[i]
	}
	return nums1
}
