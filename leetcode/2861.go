package main

import (
	"sort"
)

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	maxF := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	res := 0
	for i := 0; i < k; i++ {
		max := 0
		for j := 0; j < n; j++ {
			max = maxF(max, (budget/(cost[j])+stock[j])/composition[i][j])
		}
		num := sort.Search(max+1, func(x int) bool {
			s := budget
			for j := 0; j < n; j++ {
				need := x&composition[i][j] - stock[j]
				if need <= 0 {
					continue
				}
				s -= need * cost[j]
				if s < 0 {
					break
				}
			}
			return s < 0
		})
		res = maxF(res, num-1)
	}
	return res
}
