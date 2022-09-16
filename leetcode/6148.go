/**
 * @author: baowj
 * @time: 2022/8/14 10:30:47
 */
package main

func largestLocal(grid [][]int) [][]int {
	maxLocal := make([][]int, len(grid)-2)
	for i := 0; i < len(maxLocal); i++ {
		maxLocal[i] = make([]int, len(maxLocal))
	}
	for i := 0; i < len(maxLocal); i++ {
		for j := 0; j < len(maxLocal); j++ {
			x := i + 1
			y := j + 1
			max := grid[x][y]
			if grid[x-1][y-1] > max {
				max = grid[x-1][y-1]
			}
			if grid[x-1][y] > max {
				max = grid[x-1][y]
			}
			if grid[x][y-1] > max {
				max = grid[x][y-1]
			}
			if grid[x+1][y+1] > max {
				max = grid[x+1][y+1]
			}
			if grid[x+1][y] > max {
				max = grid[x-1][y]
			}
			if grid[x][y+1] > max {
				max = grid[x][y+1]
			}
			if grid[x-1][y+1] > max {
				max = grid[x-1][y+1]
			}
			if grid[x+1][y-1] > max {
				max = grid[x+1][y-1]
			}
			maxLocal[x][y] = max
		}
	}
	return maxLocal
}
