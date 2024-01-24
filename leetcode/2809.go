package main

import (
	"sort"
)

func minimumTime(nums1 []int, nums2 []int, x int) int {
	nums := make([][2]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		nums[i] = [2]int{nums1[i], nums2[i]}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][1] < nums[j][1]
	})

	dp, newDp := nums1, nums2
	sum1, sum2 := 0, 0
	for i := 0; i < len(dp); i++ {
		sum1 += nums1[i]
		sum2 += nums2[i]
		dp[i] = 0
	}
	if sum1 <= x {
		return 0
	}
	for i := 1; i <= len(nums); i++ {
		newDp[0] = nums[0][0] + nums[0][1]*i
		for j := 1; j < len(nums); j++ {
			newDp[j] = max(newDp[j-1], dp[j-1]+nums[j][0]+nums[j][1]*i)
		}
		sum1 += sum2
		if sum1-newDp[len(newDp)-1] <= x {
			return i
		}
		dp, newDp = newDp, dp
	}
	return -1
}
