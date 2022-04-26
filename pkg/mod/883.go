/**
 * @author: baowj
 * @time: 2022/4/26 11:00:30
 */
package main

func projectionArea(grid [][]int) int {
	xy := 0
	yGrid := make([]int, len(grid[0]))
	xGrid := make([]int, len(grid))
	for x, vec := range grid {
		for y, v := range vec {
			if v > 0 {
				xy++
			}
			if v > xGrid[x] {
				xGrid[x] = v
			}
			if v > yGrid[y] {
				yGrid[y] = v
			}
		}
	}
	for _, v := range yGrid {
		xy += v
	}
	for _, v := range xGrid {
		xy += v
	}
	return xy
}
