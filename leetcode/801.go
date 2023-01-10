/**
 * @author: baowj
 * @time: 2022/11/1 19:39:39
 */
package main

func minSwap(nums1 []int, nums2 []int) int {
	dp := make([][2]int, len(nums1))
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < len(nums1); i++ {
		if nums1[i-1] < nums1[i] && nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			if dp[i-1][0] > dp[i-1][1] {
				dp[i][0] = dp[i-1][1]
			} else {
				dp[i][0] = dp[i-1][0]
			}

			if dp[i-1][1] > dp[i-1][0] {
				dp[i][1] = dp[i-1][0] + 1
			} else {
				dp[i][1] = dp[i-1][1] + 1
			}
		} else if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		} else if nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i] {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
	}
	if dp[len(dp)-1][0] > dp[len(dp)-1][1] {
		return dp[len(dp)-1][1]
	} else {
		return dp[len(dp)-1][0]
	}
}
