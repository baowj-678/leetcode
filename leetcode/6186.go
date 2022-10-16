/**
 * @author: baowj
 * @time: 2022/9/18 10:36:33
 */
package main

func smallestSubarrays(nums []int) []int {
	dp := [32]int{}
	for i := 0; i < 32; i++ {
		dp[i] = -1
	}
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		j := 0
		for x > 0 {
			if x&1 == 1 {
				dp[j] = i
			}
			j++
			x >>= 1
		}
		max := getMax6186(dp)
		res[i] = max - i + 1
		if res[i] < 1 {
			res[i] = 1
		}
	}
	return res
}

func getMax6186(nums [32]int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
