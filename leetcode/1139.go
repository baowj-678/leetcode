/**
 * @author: baowj
 * @time: 2023/2/17 9:58
 */
package main

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowLen, colLen := make([][]int, m), make([][]int, m)
	for i := 0; i < len(grid); i++ {
		rowLen[i] = make([]int, n)
		colLen[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		length := 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				length++
			} else {
				length = 0
			}
			rowLen[i][j] = length
		}
	}

	for j := 0; j < n; j++ {
		length := 0
		for i := 0; i < m; i++ {
			if grid[i][j] == 1 {
				length++
			} else {
				length = 0
			}
			colLen[i][j] = length
		}
	}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			for l := max + 1; l <= m && l <= n; l++ {
				jl, iu := j-l+1, i-l+1
				if jl >= 0 && iu >= 0 {
					if rowLen[i][j] >= l && colLen[i][j] >= l && rowLen[iu][j] >= l && colLen[i][jl] >= l {
						max = l
					}
				} else {
					break
				}
			}
		}
	}
	return max * max
}
