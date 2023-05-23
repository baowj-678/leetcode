/**
 * @author: baowj
 * @time: 2023/5/23 10:19
 */
package main

import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	vl := make([][2]int, len(values))
	for i := 0; i < len(values); i++ {
		vl[i] = [2]int{values[i], labels[i]}
	}
	sort.Slice(vl, func(i, j int) bool {
		return vl[i][0] > vl[j][0]
	})
	m := map[int]int{}
	res := 0
	for i := 0; i < len(vl) && numWanted > 0; i++ {
		v, l := vl[i][0], vl[i][1]
		if _, ok := m[l]; !ok {
			m[l] = 0
		}
		if m[l] < useLimit {
			m[l]++
			numWanted--
			res += v
		}
	}
	return res
}
