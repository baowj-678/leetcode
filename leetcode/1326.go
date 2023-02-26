/**
 * @author: baowj
 * @time: 2023/2/25 10:26:51
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func minTaps(n int, ranges []int) int {
	newRanges := make([][2]int, len(ranges))
	for i, r := range ranges {
		newRanges[i] = [2]int{i - r, i + r}
	}
	sort.Slice(newRanges, func(i, j int) bool {
		return newRanges[i][0] < newRanges[j][0]
	})
	res := 0
	i, r := 0, 0
	for i < len(newRanges) && newRanges[i][0] <= r && r < n {
		maxR := r
		for i < len(newRanges) && newRanges[i][0] <= r {
			if newRanges[i][1] > r {
				maxR = newRanges[i][1]
			}
			i++
		}
		if maxR > r {
			res++
			r = maxR
		}
	}
	if r < n {
		return -1
	}
	return res
}
