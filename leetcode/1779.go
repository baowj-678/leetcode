/**
 * @author: baowj
 * @time: 2022/12/3 15:21:44
 * @email: bwj_678@qq.com
 */
package main

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	abs1779 := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	dis := math.MaxInt
	index := 0
	for i, point := range points {
		if x == point[0] || y == point[1] {
			tmp := abs1779(x-point[0]) + abs1779(y-point[1])
			if tmp < dis {
				dis = tmp
				index = i
			}
		}
	}
	return index
}
