/**
 * @author: baowj
 * @time: 2022/9/20 11:18:24
 */
package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum/k*k != sum {
		return false
	}
	target := sum / k
	if nums[len(nums)-1] > target {
		return false
	}
	dp := make([]int, 1<<len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i, sum := range dp {
		if sum == -1 {
			continue
		}
		for j, num := range nums {
			if (i>>j)&1 == 0 {
				next := i | (1 << j)
				if dp[next] == -1 {
					if sum+num > target {
						break
					}
					dp[next] = (sum + num) % target
				}
			}
		}
	}
	if dp[1<<len(nums)-1] == 0 {
		return true
	}
	return false
}

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	canPartitionKSubsets(nums, 4)
}
