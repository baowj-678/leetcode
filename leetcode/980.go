package main

func uniquePathsIII(grid [][]int) int {
	emptySquares := 0
	startI, startJ := 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				startI, startJ = i, j
			} else if grid[i][j] == 0 {
				emptySquares++
			}
		}
	}
	grid[startI][startJ] = 0
	return dfs980(grid, startI, startJ, 0, emptySquares)

}

func dfs980(grid [][]int, i, j, count, emptySquares int) int {
	if grid[i][j] == 2 && emptySquares == -1 {
		count++
		return count
	}
	if grid[i][j] == 0 {
		emptySquares--
		grid[i][j] = -1
		if i > 0 {
			count = dfs980(grid, i-1, j, count, emptySquares)
		}
		if i+1 < len(grid) {
			count = dfs980(grid, i+1, j, count, emptySquares)
		}
		if j > 0 {
			count = dfs980(grid, i, j-1, count, emptySquares)
		}
		if j+1 < len(grid[0]) {
			count = dfs980(grid, i, j+1, count, emptySquares)
		}
		grid[i][j] = 0
	}
	return count
}
