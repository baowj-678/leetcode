/**
 * @author: baowj
 * @time: 2023/3/30 10:08
 */
package main

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	last := points[0][0]
	res := 0
	for _, point := range points {
		if point[0]-last > res {
			res = point[0] - last
		}
		last = point[0]
	}
	return res
}
