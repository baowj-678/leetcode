package main

func distinctDifferenceArray(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums)+1)
	set := map[int]struct{}{}
	for i, num := range nums {
		set[num] = struct{}{}
		prefix[i] = len(set)
	}
	set = map[int]struct{}{}
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		set[num] = struct{}{}
		suffix[i] = len(set)
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = prefix[i] - suffix[i+1]
	}

	return res
}
