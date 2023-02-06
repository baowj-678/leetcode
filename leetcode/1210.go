/**
 * @author: baowj
 * @time: 2023/2/5 10:25:18
 * @email: bwj_678@qq.com
 */
package main

func minimumMoves(grid [][]int) int {
	n := len(grid)
	verticalDp := make([][]int, n)
	horizontalDp := make([][]int, n)
	for i := 0; i < n; i++ {
		verticalDp[i] = make([]int, n)
		horizontalDp[i] = make([]int, n)
	}

	const MAX int = 1e8
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			verticalDp[i][j] = MAX
			horizontalDp[i][j] = MAX
		}
	}
	horizontalDp[0][1] = 0
	for k := 0; k < n; k++ {
		for t := 0; t <= k; t++ {
			for i := 0; i <= k; i++ {
				j := k - i

				if grid[i][j] == 1 {
					continue
				}
				// horizontal
				// h -> left
				min := horizontalDp[i][j]
				if j-2 >= 0 && min > horizontalDp[i][j-1]+1 {
					min = horizontalDp[i][j-1] + 1
				}
				// h -> down
				if i-1 >= 0 && j-1 >= 0 && grid[i][j-1] == 0 && min > horizontalDp[i-1][j]+1 {
					min = horizontalDp[i-1][j] + 1
				}
				// v -> rotate
				if i+1 < n && j-1 >= 0 {
					if grid[i][j-1] == 0 && grid[i][j] == 0 && grid[i+1][j-1] == 0 && grid[i+1][j] == 0 {
						if min > verticalDp[i+1][j-1]+1 {
							min = verticalDp[i+1][j-1] + 1
						}
					}
				}
				horizontalDp[i][j] = min
				// vertical
				// v -> left
				min = verticalDp[i][j]
				if j-1 >= 0 && i-1 >= 0 && grid[i-1][j] == 0 && min > verticalDp[i][j-1]+1 {
					min = verticalDp[i][j-1] + 1
				}
				// v -> down
				if i-2 >= 0 && min > verticalDp[i-1][j]+1 {
					min = verticalDp[i-1][j] + 1
				}
				// v -> rotate
				if i-1 >= 0 && j+1 < n {
					if grid[i-1][j] == 0 && grid[i-1][j+1] == 0 && grid[i][j] == 0 && grid[i][j+1] == 0 {
						if min > horizontalDp[i-1][j+1]+1 {
							min = horizontalDp[i-1][j+1] + 1
						}
					}
				}
				verticalDp[i][j] = min
			}
		}
	}

	for k := n; k <= 2*n-2; k++ {
		l := 2*n - k - 1
		for t := 0; t < l; t++ {
			for i := n - 1; i >= n-l; i-- {
				j := k - i

				if grid[i][j] == 1 {
					continue
				}
				// horizontal
				// h -> left
				min := horizontalDp[i][j]
				if j-2 >= 0 && min > horizontalDp[i][j-1]+1 {
					min = horizontalDp[i][j-1] + 1
				}
				// h -> down
				if i-1 >= 0 && j-1 >= 0 && grid[i][j-1] == 0 && min > horizontalDp[i-1][j]+1 {
					min = horizontalDp[i-1][j] + 1
				}
				// v -> rotate
				if i+1 < n && j-1 >= 0 {
					if grid[i][j-1] == 0 && grid[i][j] == 0 && grid[i+1][j-1] == 0 && grid[i+1][j] == 0 {
						if min > verticalDp[i+1][j-1]+1 {
							min = verticalDp[i+1][j-1] + 1
						}
					}
				}
				horizontalDp[i][j] = min
				// vertical
				// v -> left
				min = verticalDp[i][j]
				if j-1 >= 0 && i-1 >= 0 && grid[i-1][j] == 0 && min > verticalDp[i][j-1]+1 {
					min = verticalDp[i][j-1] + 1
				}
				// v -> down
				if i-2 >= 0 && min > verticalDp[i-1][j]+1 {
					min = verticalDp[i-1][j] + 1
				}
				// v -> rotate
				if i-1 >= 0 && j+1 < n {
					if grid[i-1][j] == 0 && grid[i-1][j+1] == 0 && grid[i][j] == 0 && grid[i][j+1] == 0 {
						if min > horizontalDp[i-1][j+1]+1 {
							min = horizontalDp[i-1][j+1] + 1
						}
					}
				}
				verticalDp[i][j] = min
			}
		}
	}
	if horizontalDp[n-1][n-1] == MAX {
		return -1
	}
	return horizontalDp[n-1][n-1]
}
