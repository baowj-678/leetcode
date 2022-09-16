/**
 * @author: baowj
 * @time: 2022/9/5 10:49:14
 */
package main

import (
	"math"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] <= pairs[j][1]
	})
	cnt := 0
	maxVal := math.MinInt
	for _, pair := range pairs {
		if pair[0] > maxVal {
			maxVal = pair[1]
			cnt++
		}
	}
	return cnt
}
