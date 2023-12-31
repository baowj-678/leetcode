package main

import (
	"fmt"
	"math"
)

func minCost(nums []int, x int) int64 {
	minNums := make([]int, len(nums))
	copy(minNums, nums)
	min := func(x int, y int64) int {
		if int64(x) < y {
			return x
		}
		return int(y)
	}

	minInt64 := func(x, y int64) int64 {
		if x < y {
			return x
		}
		return y
	}

	res := int64(math.MaxInt64)

	for i := 0; i < len(nums); i++ {
		sum := int64(i * x)
		for k := 0; k < len(nums); k++ {
			minNums[k] = min(minNums[k], int64(nums[(k-i+len(nums))%len(nums)]))
			sum += int64(minNums[k])
		}
		res = minInt64(res, sum)
	}
	return res
}

func main() {
	nums := []int{20, 1, 15}
	fmt.Println(minCost(nums, 5))
}
