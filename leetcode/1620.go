/**
 * @author: baowj
 * @time: 2022/11/2 15:01:53
 */
package main

import "math"

func bestCoordinate(towers [][]int, radius int) []int {
	resI, resJ := 0, 0
	var maxSignal float64 = 0
	for i := 0; i <= 50; i++ {
		for j := 0; j <= 50; j++ {
			var sig float64 = 0
			for _, tower := range towers {
				sig += computeSignal(i, j, tower[0], tower[1], radius, tower[2])
			}
			if sig > maxSignal {
				maxSignal = sig
				resI = i
				resJ = j
			}
		}
	}
	return []int{resI, resJ}
}

func computeSignal(i, j, x, y, r, q int) float64 {
	square := (i-x)*(i-x) + (j-y)*(j-y)
	if square > r*r {
		return 0
	}
	d := math.Sqrt(float64(square))
	return math.Floor(float64(q) / (1 + d))
}
