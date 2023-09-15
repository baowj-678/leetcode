/**
 * @author: baowj
 * @time: 2023/9/14 10:15
 */
package main

import "sort"

func checkValidGrid(grid [][]int) bool {
	n := len(grid)
	positions := make([][3]int, 0, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			positions = append(positions, [3]int{grid[i][j], i, j})
		}
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i][0] < positions[j][0]
	})

	absInt := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}
	check := func(a, b []int) bool {
		x := absInt(a[0] - b[0])
		y := absInt(a[1] - b[1])
		if (x == 2 && y == 1) || (x == 1 && y == 2) {
			return true
		}
		return false
	}
	tmp := positions[0][1:]
	if tmp[0] != 0 || tmp[1] != 0 {
		return false
	}
	for _, p := range positions[1:] {
		if !check(tmp, p[1:]) {
			return false
		}
		tmp[0], tmp[1] = p[1], p[2]
	}
	return true
}

func main() {
	grid := [][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}}
	checkValidGrid(grid)
}
