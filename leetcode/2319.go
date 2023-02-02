/**
 * @author: baowj
 * @time: 2023/2/2 21:42:21
 * @email: bwj_678@qq.com
 */
package main

func checkXMatrix(grid [][]int) bool {
	m := len(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i == j || i+j == m-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else if grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
