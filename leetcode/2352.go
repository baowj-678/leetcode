/**
 * @author: baowj
 * @time: 2023/6/6 9:59
 */
package main

func equalPairs(grid [][]int) int {
	n := len(grid)
	check := func(x, y int) bool {
		for i := 0; i < n; i++ {
			if grid[x][i] != grid[i][y] {
				return false
			}
		}
		return true
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if check(i, j) {
				res++
			}
		}
	}
	return res
}
