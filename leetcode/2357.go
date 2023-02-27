/**
 * @author: baowj
 * @time: 2023/2/24 15:52
 */
package main

import "sort"

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	m := map[int]struct{}{}
	for _, num := range nums {
		m[num] = struct{}{}
	}
	res := 0
	for k, _ := range m {
		if k != 0 {
			res++
		}
	}
	return res
}
