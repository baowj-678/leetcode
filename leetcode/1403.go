/**
 * @author: baowj
 * @time: 2022/8/5 10:32:02
 */
package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sums := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sums[i] = sum
	}
	for i := 0; i < len(nums); i++ {
		if sums[i] > sum-sums[i] {
			return nums[:i+1]
		}
	}
	return nums
}

func main() {
	a := []int{2, 3, 6, 1, 2, 3, 56, 90}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a)
}
