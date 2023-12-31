package main

import "sort"

func tupleSameProduct(nums []int) int {
	cnt := map[int]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			s := nums[i] * nums[j]
			if v, ok := cnt[s]; ok {
				cnt[s] = v + 1
			} else {
				cnt[s] = 1
			}
		}
	}
	res := 0
	for _, v := range cnt {
		res += v * (v - 1) * 4
	}
	return res
}
