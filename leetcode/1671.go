package main

import (
	"fmt"
)

func minimumMountainRemovals(nums []int) int {
	res := 0
	leftUp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		leftUp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && leftUp[i] < leftUp[j]+1 {
				leftUp[i] = leftUp[j] + 1
			}
		}
	}
	rightUp := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		rightUp[i] = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] && rightUp[j]+1 > rightUp[i] {
				rightUp[i] = rightUp[j] + 1
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if leftUp[i] > 1 && rightUp[i] > 1 && leftUp[i]+rightUp[i]-1 > res {
			res = leftUp[i] + rightUp[i] - 1
		}
	}
	return len(nums) - res
}

func main() {
	nums := []int{23, 47, 63, 72, 81, 99, 88, 55, 21, 33, 32}
	fmt.Println(minimumMountainRemovals(nums))
}
