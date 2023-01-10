package main

func subarrayLCM(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		x := 1
		for j := i; j < len(nums); j++ {
			x = x * nums[j] / lcd6234(x, nums[j])
			if x == k {
				res++
			} else if x > k {
				break
			}
		}
	}
	return res
}

func lcd6234(x int, y int) int {
	if x < y {
		x, y = y, x
	}
	for x != y && y != 0 {
		x, y = y, x%y
	}
	return x
}
